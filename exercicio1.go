package main

import "fmt"

func main() {
	var a = 3
	fmt.Println(a)
	somaprimerotermos(&a)
	fmt.Println(a)
}

func somaprimerotermos(ptr *int) {
	var i int
	var soma int
	for i = 1; i <= *ptr; i++ {
		soma += i
	}
	*ptr = soma
}
