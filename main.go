package main

import (
        "fmt"
		"github.com/richraid21/go-data-structures/structures"
)

func main() {
    s := structures.NewStack()
	s.Push(&structures.Node{Value: 17})
	s.Push(&structures.Node{Value: 72})
	s.Push(&structures.Node{Value: 2})
	fmt.Println(s.Depth())
}