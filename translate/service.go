package translate

import "fvpFormatTranslator/validate"


// Entrada principal
type ServiceTranslator interface {
	// Esta funcion separa intereses
	// Valida formatos y si es correcto
	// Traduce el texto
	Translate(textoATraducir, formatoOrigen, formatoDestino string) (string, error)
}

type serviceImp struct {
	validator validate.Validator
	validatorDestiono validate.Validator
	// Dependencia para traducir texto a binario
	textToBinary Translator
}

func NewServiceImp(validator validate.Validator, validatorDestiono validate.Validator, textToBinary Translator) *serviceImp  {
	return &serviceImp{
		validator: validator,
		validatorDestiono: validatorDestiono,
		textToBinary: textToBinary,
	}
}

func (si *serviceImp) Translate(textoATraducir, formatoOrigen, formatoDestino string) (string, error) {

	// Validamos el formato de origen
	if err := si.validator.Validate(formatoOrigen) ; err != nil {
		return "", err
	}

	// Validamos formato destino
	if err := si.validatorDestiono.Validate(formatoDestino) ; err != nil {
		return "", err
	}

	switch formatoOrigen {
	case "TEXT":
		switch formatoDestino {
		case "BINARY":
			return si.textToBinary.Translate(textoATraducir)

			// Case MORSE

			// Case otro formato
		}

		// Case Binary

		// Case MORSE

	}

	// Resultado exitoso
	return "test", nil
}