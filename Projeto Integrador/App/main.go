package main

import (
	"log"
	"net/http"
)

func main () {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	porta := ":8080"
	log.Println("Servidor rodando em http://localhost" + porta)
	err := http.ListenAndServe(porta, nil)
	
	if err != nil {
		log.Fatal(err)
	} 
}
