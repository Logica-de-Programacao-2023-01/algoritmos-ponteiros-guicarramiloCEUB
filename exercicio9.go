package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
	preco  float64
}

func main() {
	l := Livro{titulo: "Gui", autor: "Gui", preco: 59.9}
	fmt.Println(l.titulo)
	fmt.Println(l.autor)
	fmt.Println(l.preco)
	dezpor(&l)
	fmt.Println(l.titulo)
	fmt.Println(l.autor)
	fmt.Println(l.preco)

}

func dezpor(ptr *Livro) {
	(*ptr).preco *= 1.1
}
