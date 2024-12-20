package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IpDatas struct {
	IP         string `json:"query"`
	Pais       string `json:"country"`
	CodigoPais string `json:"countryCode"`
	Estado     string `json:"regionName"`
	Municipio  string `json:"city"`
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

	return "\nfim"

}
