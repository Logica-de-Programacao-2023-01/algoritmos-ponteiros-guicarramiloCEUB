package main

import "fmt"

func main() {
	var a = 4.6
	fmt.Println(a)
	mediapi(&a)
	fmt.Println(a)
}

func mediapi(ptr *float64) {
	*ptr = (*ptr + 3.14) / 2
}
