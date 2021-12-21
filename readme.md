# config-parser

A parser for a very simple type of config. Returns a tree where `Value` is the value at that point in the node, and `Children` are the children of the tree. Members of the tree do not know their own key, rather it is handled by the parent of the tree.

An example file this could parse -
```
# a comment, comments must be aligned with the left for now
food:
  fruits:
# you can just put in nothing for the key to create a list
    apples:
    bananas:
    pears: good
# pears are good
  vegetables:
    broccoli:
    corn:
    brussel sprouts: bad
```

Please work...
