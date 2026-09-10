package pointers

import "fmt"

type Pessoa struct {
	Nome string
}

func People() {
	var p1 Pessoa = Pessoa{Nome: "Giovani"}
	fmt.Println("-----PESSOAS------")
	fmt.Println("Antes da modificação: ", p1.Nome)
	var p3 *Pessoa = &p1
	fmt.Println(p3.Nome)
	p3.Nome = "João"
	fmt.Println("Depois da modificação:", p1.Nome)
}