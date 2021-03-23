package pkg

/*

	Stacks implementation in Go
	Uses singly linked lists principle to maintain efficient time complexity


	To create a stack: s := Stack{}

*/

/*
*	holds the node of a stack list
 */
type Node struct {
	Value interface{}
	Next  *Node
}

/*
*	the actual stack
 */
type Stack struct {
	Length uint
	Head   *Node
	Tail   *Node
}

/*
*	Push element onto stack
	O(1)
*/
func (s *Stack) Add(v interface{}) *Stack {
	// make node
	n := &Node{Value: v}

	// check if head is nil
	if s.Head == nil {
		s.Head = n
		s.Tail = s.Head
	} else

	// otherwise bind node's next to tail and tail to node
	{
		n.Next = s.Tail
		s.Tail = n
	}

	s.Length++

	return s
}

/*
*	Retrieve the last pushed element and delete it from stack
	O(1)
*/
func (s *Stack) Pop() *Node {
	// check length we may not need to return anything :)
	if s.Length == 0 {
		return nil
	} else {
		n := s.Tail
		s.Tail = s.Tail.Next
		n.Next = nil

		if s.Length == 1 {
			s.Head = nil
		}

		s.Length--

		return n
	}
}

/*
*	Loop the stack with a callback
	O(n)
*/
func (s *Stack) Traverse(f func(*Node)) *Stack {
	n := s.Tail

	for n != nil {
		f(n)
		n = n.Next
	}

	return s
}

/*
*	Find a node containing via callback
	Lookup is done from tail to head
	Callback receives a node as parameter and must return bool

*	O(n)
*/
func (s *Stack) Find(f func(*Node) bool) *Node {
	n := s.Tail

	for n != nil {
		if f(n) == true {
			return n
		}
		n = n.Next
	}

	return nil
}
