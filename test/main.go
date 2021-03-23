package main

import (
	"fmt"

	Stacks "github.com/tgbv/go-stacks/pkg"
)

func main() {
	s := Stacks.Stack{}

	// addition & traversion
	s.Add("hello").Add("meme").Add("hero").Traverse(func(n *Stacks.Node) {
		fmt.Println(n)
	})

	// pop
	fmt.Println(s.Pop())

	// find
	fmt.Println(s.Find(func(n *Stacks.Node) bool {
		if n.Value == "meme" {
			return true
		} else {
			return false
		}
	}))

}
