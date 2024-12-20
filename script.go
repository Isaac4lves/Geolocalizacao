package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type IpDatas struct {
	IP         string  `json:"query"`
	Pais       string  `json:"country"`
	CodigoPais string  `json:"countryCode"`
	Estado     string  `json:"regionName"`
	Municipio  string  `json:"city"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"lon"`
}

func RequestApi(ip string, tipo string) string {
	var url = "http://ip-api.com/json/" + ip

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "\nfim"
	}
	defer response.Body.Close()

	var dados IpDatas
	if err := json.NewDecoder(response.Body).Decode(&dados); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return "\nfim"
	}

	if tipo == "ip" {
		return dados.IP
	}

	if tipo == "estado" {
		return dados.Estado
	}
	if tipo == "municipio" {
		return dados.Municipio
	}
	if tipo == "pais" {
		return dados.Pais
	}
	if tipo == "codigoPais" {
		return dados.CodigoPais
	}
	if tipo == "latitude" {
		// Aqui, utilizamos a função strconv.FormatFloat, que recebe quatro parâmetros:
		// o número a ser convertido (valor), o formato ('f' para ponto fixo), a precisão
		// (-1 para usar a precisão mínima necessária para representar o valor sem perda de dados)
		//  e o tamanho em bits do número (64 para especificar que o número é do tipo float64).
		//  Essa função retorna o valor formatado como uma string, que é então armazenado na variável strValor.
		lat := strconv.FormatFloat(dados.Latitude, 'f', -1, 64)
		return lat
	}
	if tipo == "longitude" {
		lon := strconv.FormatFloat(dados.Longitude, 'f', -1, 64)
		return lon
	}
	return "\nfim"

}

func Getip() string {
	response, err := http.Get("https://meuip.com/api/meuip.php")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
	}
	ip := string(body)
	return ip
}
