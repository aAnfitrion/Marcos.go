package main

import (
	"fmt"

	"github.com/aAnfitrion/Marcos.go"
	"github.com/aAnfitrion/Marcos.go/bordes"
)

func main() {
	menu := marcos.Marcos{}
	menu.SetTitulo("Menu de inicio")
	menu.SetContenidoSlice([]string{
		"1) Hola",
		"2) Mundo",
		"3) Epic",
	})
	menu.SetContenidoString("1) Hola\n2) Mundo\n3) Epic")
	menu.Enmarcar(bordes.Recto)

	fmt.Println(menu.GetMarcoString())
}
