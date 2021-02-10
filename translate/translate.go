package translate

import "fmt"

// Interfaz para implemnetar patron estrategia
// Cada unto de los algoritmos de conversion deberan de implementar esta interfaz
type Translator interface {
	Translate(val string) (string, error)
}

// Struct para Implementacion de texto a binario
type textToBinaryImpl struct {
}

func NewtextToBinaryImpl() *textToBinaryImpl {
	return &textToBinaryImpl{}
}

// Algoritmo para traducir texto a binario
func (ttob *textToBinaryImpl) Translate(val string) (string, error) {

	response := ""
	// traducimos cada caracter a binario y lo adicionamos a la respuesta
	for _, c := range val {
		response = fmt.Sprintf("%s%.8b", response, c)
	}

	return response, nil
}

// Struct para Implementacion de algoritmo de binario a texto
type binaryToTextImpl struct {
}

func NewBinaryToTextImpl() *binaryToTextImpl {
	return &binaryToTextImpl{}
}

// Algoritmo para traducir texto a binario
func (btoText *binaryToTextImpl) Translate(val string) (string, error) {

	response := ""

	// Realizar la traduccion

	return response, nil
}
