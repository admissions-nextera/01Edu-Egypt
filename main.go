package main

import (
	"log"
)

type LargeStruct struct {
	data [100000000]int // 8kb
}

func processValue(s LargeStruct) int {
	return s.data[0]
}

func processPointer(s *LargeStruct) int {
	return (s.data[0])
}

func main() {
	large := LargeStruct{}
	log.Println(processValue(large))  // copy all 8kb
	log.Print(processPointer(&large)) // copy address
}
