package configparser

import (
	"bufio"
	"os"
)

type ConfigTree struct { // Basic tree
	Value    string // Its a string for now
	Children map[string]ConfigTree
}

type TreeStack struct { // A simple linked-list stack implementation
	Value    *ConfigTree // Holds a tree
	previous *StrStack   // And the previous stack
}

func (s *TreeStack) Push(val string) (ret *TreeStack) {
	ret = new(TreeStack)
	ret.previous = s
	return ret
}

func (s *TreeStack) Pop() TreeStack {
	return s.previous // feels wrong not to deallocate the top of the stack but oh well
}

func parseConfigLine(line string, prev_indent, indent_base int) (key, value string, indent_change int) {
	i := 0
	for line[i] == " " || line[i] == "\t" { // Get up to stop of whitespace
		i++
	}
	indent_change = i - prev_indent // Measure whitespace and see how it changed between first and last run
	j := i
	for line[j] != ":" && line[j-1] == "\\" { // Get key
		j++
	}
	key = line[i:j]
	for (line[j] != " " && line[j] != "\t") && j < len(line) { // Get value
		j++
	}
	if line[j] == " " || line[j] == "\t" { // If value is not found, return a null string
		value = "\000"
	} else {
		value = line[j:]
	}
	return
}

func ParseConfigLines(lines []string) ConfigTree {
	ret := ConfigTree{Value: "\000", Children: make(map[string]ConfigTree)}
	iter := *TreeStack{Value: &ret, previous: nil}
	indent, indent_base := 0
	whitespace_finder := 0
	for _, line := range lines {
		if line[0] == "#" { // Comment checking
			continue
		}

		for whitespace_finder = 0; (line[whitespace_finder] == " " || line[whitespace_finder] == "\t") && whitespace_finder < len(line); whitespace_finder++ {
		} // Make sure no empty lines exist
		if whitespace_finder == len(line) {
			continue
		}

		if indent_base == 0 && (line[0] == " " || line[0] == "\t") {
			for ; line[indent_base] == " " || line[indent_base] == "\t"; indent_base++ {
			} // Increase the indent base for each whitespace found, the first time indents are found
		}

		key, value, indent_change := parseConfigLine(line, indent) // Parse line
		indent += indent_change                                    // Change the indent

		iter.Value.Children[key] = ConfigTree{Value: value, make(map[string]ConfigTree)} // Add a new child to the current top of stack
		iter = iter.Push(iter.Value.Children[key])                                       // Push the new child onto the stack

		if indent_change < 1 { // Pop the stack for each negative change in indent
			iter.Pop()
			if indent_base != 0 {
				for i := indent_change / indent_base; i < 0; i++ {
					iter = iter.Pop()
				}
			}
		}
	}
	return ret
}

func ParseConfig(filename string) ConfigTree { // All this script does is read the file line by line, then pipe it through ParseConfigLines().
	file, err = os.Open(filename)
	if err != nil {
		return ConfigTree{Value: string(err)}
	}

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return ParseConfigLines(lines)
}
