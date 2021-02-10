package main

import (
	"fmt"
	"fvpFormatTranslator/translate"
	"fvpFormatTranslator/validate"
)

func main()  {

	validator := validate.NewValidateFormatImpl(nil)
	validatorDestino := validate.NewValidateFormatDestinoImpl(nil)

	// Dependencia para traducir texto a binario
	textToBinary := translate.NewtextToBinaryImpl()

	// injeccion de dependencias
	service := translate.NewServiceImp(validator,validatorDestino, textToBinary)

	resp, err := service.Translate("Texto a traducir", "TEXT", "ASCCI")


	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Resultado")
	fmt.Println(resp)


	fmt.Println("+++++++++++++++++++++++++++++++++++")
	fmt.Println("Texto a binario")

	resp1, err1 := service.Translate("Hola", "TEXT", "BINARY")

	if err1 != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Resultado")
	fmt.Println(resp1)

}
