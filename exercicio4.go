package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 1234
	fmt.Println(a)
	somaprimeiros2(&a)
	fmt.Println(a)

}

func somaprimeiros2(ptr *int) {
	var str = strconv.Itoa(*ptr)
	var i int
	soma := 0
	if len(str) >= 2 {
		for i = len(str) - 1; i >= len(str)-2; i++ {
			num, _ := strconv.Atoi(string(str[i]))
			soma += num
		}
	}
	*ptr = soma
}
