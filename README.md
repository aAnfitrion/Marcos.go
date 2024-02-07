# Marcos.go
Este paquete sirve para añadir un marco/borde a un string o slice de strings.
## Modo de uso
Aquí un ejemplo de como usar Marcos.go
```go
package main

import (
	"fmt"

	"github.com/aAnfitrion/Marcos.go"
	"github.com/aAnfitrion/Marcos.go/bordes"
)

func main() {
	menu := marcos.Marcos{}
	// El título es opcional
	menu.SetTitulo("Menu de inicio")
	menu.SetContenidoSlice([]string{
		"1) Hola",
		"2) Mundo",
		"3) Epic",
	})
	menu.Enmarcar(bordes.Recto)

	fmt.Println(menu.GetMarcoString())
}
```
Esto nos devolverá:
```
┌─ Menu de inicio ─┐
│ 1) Hola          │
│ 2) Mundo         │
│ 3) Epic          │
└──────────────────┘
```
Podemos pasar un string con saltos de linea en lugar de un slice de strings usando el método `SetContenidoString(string)`:
```go
menu.SetContenidoString("1) Hola\n2) Mundo\n3) Epic")
```
Puedes usar tus propios bordes, solo tienes que seguir la siguiente estructura de `Marcos.go/bordes/tipo.go`:
```go
package bordes

type Bordes struct {
	Horizontal, Vertical, EsquinaA, EsquinaB, EsquinaC, EsquinaD string
}
```
Este paquete ya viene con tres conjuntos de bordes predefinidos:
```go
var Recto Bordes = Bordes{
	Horizontal: "─",
	Vertical:   "│",
	EsquinaA:   "┌",
	EsquinaB:   "┐",
	EsquinaC:   "└",
	EsquinaD:   "┘",
}

var Redondeado Bordes = Bordes{
	Horizontal: "─",
	Vertical:   "│",
	EsquinaA:   "╭",
	EsquinaB:   "╮",
	EsquinaC:   "╰",
	EsquinaD:   "╯",
}

var Doble Bordes = Bordes{
	Horizontal: "═",
	Vertical:   "║",
	EsquinaA:   "╔",
	EsquinaB:   "╗",
	EsquinaC:   "╚",
	EsquinaD:   "╝",
}
```
