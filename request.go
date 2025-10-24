package main

import (
	"aler93.com/adressongo/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	_ "strconv"
	"time"
)

func viaCep(cep string) {
	// call https://viacep.com.br/ws/97020100/json
	var url string = "https://viacep.com.br/ws/" + cep + "/json"
	resp, err := http.Get(url)
	if err != nil {
		logErrorToDisk(err.Error())
		os.Exit(2)
	}

	response, errIo := io.ReadAll(resp.Body)
	if errIo != nil {
		panic(errIo)
	}
	resp.Body.Close()

	var tmp model.ViaCep
	json.Unmarshal(response, &tmp)

	if len(result.Cep) <= 0 {
		result.Cep = tmp.Cep
		result.Rua = tmp.Logradouro
		result.Bairro = tmp.Bairro
		result.Cidade = tmp.Localidade
		result.Uf = tmp.Uf
		result.Unidade = tmp.Unidade
		result.Regiao = tmp.Regiao
		result.Complemento = tmp.Complemento
		result.Ibge = tmp.Ibge
		result.Ddd = tmp.Ddd
		result.Siafi = tmp.Siafi
		result.Origem = "ViaCep"
	} else {
		// Check update
		//var doUpdate bool = false
	}
}

func openCep(cep string) {
	// call https://opencep.com/v1/97020100
	var url string = "https://opencep.com/v1/" + cep
	resp, err := http.Get(url)
	if err != nil {
		logErrorToDisk(err.Error())
		os.Exit(2)
	}

	response, errIo := io.ReadAll(resp.Body)
	if errIo != nil {
		panic(errIo)
	}
	resp.Body.Close()

	var tmp model.OpenCep
	json.Unmarshal(response, &tmp)

	if len(result.Cep) <= 0 {
		result.Cep = tmp.Cep
		result.Rua = tmp.Logradouro
		result.Bairro = tmp.Bairro
		result.Cidade = tmp.Localidade
		result.Uf = tmp.Uf
		result.Complemento = tmp.Complemento
		result.Origem = "OpenCep"
	} else {
		// Check update
		//var doUpdate bool = false
	}
}

func brasilApi(cep string) {
	// call https://brasilapi.com.br/api/cep/v2/97020100
	var url string = "https://brasilapi.com.br/api/cep/v2/" + cep
	resp, err := http.Get(url)
	if err != nil {
		logErrorToDisk(err.Error())
		os.Exit(2)
	}

	response, errIo := io.ReadAll(resp.Body)
	if errIo != nil {
		panic(errIo)
	}
	resp.Body.Close()

	var tmp model.BrasilApi
	json.Unmarshal(response, &tmp)

	if len(result.Cep) <= 0 {
		result.Cep = tmp.Cep
		result.Rua = tmp.Street
		result.Bairro = tmp.Neighborhood
		result.Cidade = tmp.City
		result.Uf = tmp.State
		result.Origem = "BrasilApi"
	} else {
		// Check update
		//var doUpdate bool = false
	}
}

func logErrorToDisk(message string) {
	now := time.Now()
	year := fmt.Sprintf("%d", now.Year())
	month := fmt.Sprintf("%d", now.Month())
	day := fmt.Sprintf("%d", now.Day())
	hour := fmt.Sprintf("%d", now.Hour())
	min := fmt.Sprintf("%d", now.Minute())
	sec := fmt.Sprintf("%d", now.Second())

	//fName := "request_error_" + year + month + day + ".log"
	fName := "request_error_" + year + month + day + hour + min + sec + ".log"
	text := year + "-" + month + "-" + day + " " + hour + ":" + min + ":" + sec + "\n - " + message
	//text := now.Format("2000-12-30_23-59-58") + "\n - " + message

	d1 := []byte(text)
	err := os.WriteFile(App.Log+fName, d1, 0644)

	if err != nil {
		panic(err)
	}
}
