package main

import "fmt"

func main() {
	a := Stack{
		data: []int{1, 2, 3, 4, 5, 6},
	}
	b := Stack{
		data: []int{1, 2, 3},
	}
	rra(&a)
	fmt.Println(a.data)
	fmt.Println(b.data)
}
