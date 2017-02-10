package main

import (
        "fmt"
		"github.com/richraid21/go-data-structures/structures"
)

func main() {
    s := structures.NewStack()
	s.Push(17)
	s.Push(72)
	s.Push(2)
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	fmt.Println(s.Empty())
}