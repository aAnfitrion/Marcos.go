package marcos

import (
	"strings"

	"github.com/aAnfitrion/Marcos.go/bordes"
)

type Marcos struct {
	titulo, marco string
	marcoSlice    []string
}

func (m *Marcos) SetTitulo(titulo string) {
	m.titulo = titulo
}

func (m *Marcos) SetContenidoString(contenido string) {
	m.marco = contenido
	m.marcoSlice = strings.Split(contenido, "\n")
}

func (m *Marcos) SetContenidoSlice(contenido []string) {
	m.marcoSlice = contenido
	m.marco = strings.Join(contenido, "\n")
}

func (m *Marcos) Enmarcar(bordes bordes.Bordes) {
	var stringMasLargo int

	for _, linea := range m.marcoSlice {
		if len(linea) < stringMasLargo {
			continue
		}
		stringMasLargo = len(linea)
	}

	var capa string

	if m.titulo == "" {
		capa = bordes.EsquinaA +
			strings.Repeat(bordes.Horizontal, stringMasLargo+2) +
			bordes.EsquinaB
	} else {
		if len(m.titulo) > stringMasLargo {
			stringMasLargo = len(m.titulo) + 2
		}

		horizontales := strings.Repeat(
			bordes.Horizontal, stringMasLargo-len(m.titulo)-2)

		capa = bordes.EsquinaA + bordes.Horizontal + " " + m.titulo +
			" " + bordes.Horizontal + horizontales + bordes.EsquinaB
	}

	var cuerpo []string
	for _, linea := range m.marcoSlice {
		espaciosExtra := strings.Repeat(" ", stringMasLargo-len(linea))

		cuerpo = append(cuerpo,
			bordes.Vertical+" "+linea+espaciosExtra+" "+bordes.Vertical)
	}

	base := bordes.EsquinaC +
		strings.Repeat(bordes.Horizontal, stringMasLargo+2) +
		bordes.EsquinaD

	contenido := []string{capa}
	contenido = append(contenido, cuerpo...)
	contenido = append(contenido, base)

	m.SetContenidoSlice(contenido)
}

func (m *Marcos) GetMarcoString() string {
	return m.marco
}
