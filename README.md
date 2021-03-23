# go-stacks

Go stacks implemented from scratch.
Uses singly linked lists principle to maintain efficient time complexity

To create a stack: 
```
import s "github.com/tgbv/go-stacks/pkg"

v := s.Stack{}

// or

var v s.Stack = s.Stack{}
```

Available methods:
```
func (s *Stack) Add(v interface{}) *Stack

func (s *Stack) Pop() *Node

func (s *Stack) Traverse(f func(*Node)) *Stack

func (s *Stack) Find(f func(*Node) bool) *Node
```
