package main

import (
	"container/list"
	"fmt"
)

func main() {
	sayhello()
	fmt.Println(substr("Hello!", 2, 4))
	fmt.Println(int('H' - 'A'))
	fmt.Println(len("sss"))
	var l list.List
	l.PushBack("1")
	l.PushBack("2")
	fmt.Println(l.Len())
	fmt.Println(reverse_string("Hello!"))
	fmt.Println(test_two(1, 2))
}
