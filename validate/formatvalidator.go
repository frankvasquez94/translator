package validate

import (
	"errors"
	"strings"
)

const validFormats = "TEXT, BINARY, MORSE"
const validFormatsDestino = "TEXT, BINARY, MORSE"


// Interfaz para implementar el patron cadena de responsabilidad
type Validator interface {
	Validate(format string) error
}


type validateFormatImpl struct {
	// siguiente step en la cadena de responsabilidad.
	next Validator
}

func NewValidateFormatImpl(next Validator) *validateFormatImpl {
	return &validateFormatImpl{
		next: next,
	}
}


func (vfi *validateFormatImpl) Validate(format string) error {

	// Si no contiene el formato, devuelve error
	if !strings.Contains(validFormats, format) {
		return errors.New("Formato Invalido")
	}

	// si lo contiene verificamos si la cadena de responsabilidad continua

	if vfi.next != nil {
		vfi.next.Validate(format)
	}

	// Si next en la cadena de responsabilidad es nil
	// devolvemos nil, indicando que el resultado es exitoso
	return nil

}


type validateFormatDestinoImpl struct {
	// siguiente step en la cadena de responsabilidad.
	next Validator
}

func NewValidateFormatDestinoImpl(next Validator) *validateFormatImpl {
	return &validateFormatImpl{
		next: next,
	}
}


func (vfi *validateFormatDestinoImpl) Validate(format string) error {

	// Si no contiene el formato, devuelve error
	if !strings.Contains(validFormatsDestino, format) {
		return errors.New("Formato Invalido")
	}

	// si lo contiene verificamos si la cadena de responsabilidad continua

	if vfi.next != nil {
		vfi.next.Validate(format)
	}

	// Si next en la cadena de responsabilidad es nil
	// devolvemos nil, indicando que el resultado es exitoso
	return nil

}
