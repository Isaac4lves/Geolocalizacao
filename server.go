package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type HomeData struct {
	IP         string
	Pais       string
	Codigopais string
	Estado     string
	Municipio  string
	Latitude   string
	Longitude  string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	ip := Getip()
	data := HomeData{
		IP:         RequestApi(ip, "ip"),
		Pais:       RequestApi(ip, "pais"),
		Codigopais: RequestApi(ip, "codigoPais"),
		Estado:     RequestApi(ip, "estado"),
		Municipio:  RequestApi(ip, "municipio"),
		Latitude:   RequestApi(ip, "latitude"),
		Longitude:  RequestApi(ip, "longitude"),
	}

	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homePage)

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
