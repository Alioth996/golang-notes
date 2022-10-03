package main

import "fmt"

func main() {
	hi := "hello world"
	fmt.Println(hi)

	fmt.Printf("isEven(5): %v\n", isEven(4))
}

func isEven(num int) bool {
	return num%2 == 0
}
