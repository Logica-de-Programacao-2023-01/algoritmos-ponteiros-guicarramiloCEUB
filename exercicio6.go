package main

import "fmt"

type livro struct {
	titulo string
	autor  string
}

func main() {
	l := livro{titulo: "aaa", autor: "bbbb"}
	fmt.Println(l.autor)
	fmt.Println(l.titulo)
	alterarlivro(&l)
	fmt.Println(l.autor)
	fmt.Println(l.titulo)
}

func alterarlivro(ptr *livro) {
	(*ptr).titulo = "Desconhecido"
	(*ptr).autor = "Anonimo"
}
