package main

import "fmt"

func main() {
	num := 4
	fmt.Println(num)
	troca(&num)
	fmt.Println(num)

}

func troca(ptr *int) {
	*ptr = 0
}