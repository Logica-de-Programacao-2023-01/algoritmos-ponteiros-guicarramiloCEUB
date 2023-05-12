package main

import "fmt"

type conta struct {
	saldo   float64
	titular string
}

func main() {
	c := conta{titular: "", saldo: 0.0}
	fmt.Println(c.titular)
	fmt.Println(c.saldo)
	alterardados(&c)
	fmt.Println(c.titular)
	fmt.Println(c.saldo)

}

func alterardados(ptr *conta) {
	var saldo float64
	var titular string

	(*ptr).titular = titular
	(*ptr).saldo = saldo
}
