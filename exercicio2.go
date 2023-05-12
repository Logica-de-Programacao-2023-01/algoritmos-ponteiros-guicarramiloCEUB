package main

import "fmt"

func main() {
	var a = 3
	fmt.Println(a)
	parouimpar(&a)
	fmt.Println(a)

}

func parouimpar(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
}
