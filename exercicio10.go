package main

import "fmt"

func main() {
	slice := []int{}
	n := 4
	encherslice(&slice, n)
	fmt.Println(slice)
}

func encherslice(ptr *[]int, n int) {
	num := 0
	for i := 0; i < n; i++ {
		fmt.Println("insira numero: ")
		fmt.Scanln(&num)
		*ptr = append(*ptr, num)
	}
}
