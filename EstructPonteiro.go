package main

import "fmt"

type A struct {
	nome string
}
type B struct {
	nome string
}

func main() {
	a := A{}
	b := B{}
	var p_a *string
	var p_b *string
	p_a = &a.nome
	p_b = &b.nome
	var vetor = [2]*string{p_a, p_b}
	for i := 0; i < 2; i++ {
		fmt.Print("digite um nome:")
		fmt.Scan(vetor[i])
	}
	fmt.Println(a.nome)
	fmt.Print(b.nome)
}
