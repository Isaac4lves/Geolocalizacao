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
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")

	data := HomeData{
		IP:         RequestApi("[seu ip]", "ip"),
		Pais:       RequestApi("[seu ip]", "pais"),
		Codigopais: RequestApi("[seu ip]", "codigoPais"),
		Estado:     RequestApi("[seu ip]", "estado"),
		Municipio:  RequestApi("[seu ip]", "municipio"),
	}

	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homePage)

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
