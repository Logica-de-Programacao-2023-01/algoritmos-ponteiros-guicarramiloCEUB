package main

import "fmt"

func main() {
	var str = "ceub"
	fmt.Println(str)
	inverterstr(&str)
	fmt.Println(str)
}

func inverterstr(ptr *string) {
	var i int
	var reverse = ""
	for i = len(*ptr) - 1; i >= 0; i-- {
		fmt.Println(i)
		reverse = reverse + string((*ptr)[i])
	}
	*ptr = reverse
}
