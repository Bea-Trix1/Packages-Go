package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"` // Aqui eu defino o nome do campo no JSON
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	res, err := json.Marshal(conta) // Aqui eu converto a minha struct para um JSON utilizando o pacote encoding/json - marshal
	if err != nil {
		panic(err)
	}
	println(string(res)) // E printo o JSON

	err = json.NewEncoder(os.Stdout).Encode(conta) // Aqui eu crio um encoder para printar o JSON
	if err != nil {
		panic(err)
	}
	// Quando eu uso Marshal, eu guardo o valor para mim, e quando eu uso o NewEncoder, já encodo o valor para o destino que eu quero

	var conta2 Conta
	err = json.Unmarshal(res, &conta2) // Aqui eu converto o JSON para a struct utilizando o pacote encoding/json - unmarshal
	// Ele aponta para o endereço de memória da struct, e assim ele consegue alterar o valor da struct
	if err != nil {
		panic(err)
	}

	println(conta2.Saldo)
}
