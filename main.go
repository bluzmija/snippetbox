package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Definicija home handler funkcije koja ispisuje byte slice
// koji sadrzi "Hello from Snipperbox"
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	//Izvadi vrijednost od ID wildcarda pomoci r.PathValue i pokusaj ga
	//pretvoriti u integer koristeci strconv.Atoi, ako ga nepres pretvoriti
	//ili je id < 1 -- izbaci error i 404
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a snippet"))
}

func main() {

	//koristimo http.NewServeMux da inicijaliziramo novi servemux
	mux := http.NewServeMux()
	//ove funkcije "bindaju" odredenu funckiju na odredeni URL pattern
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	//u konzolu isprintaj log da server starta
	log.Print("starting server on :4000")

	//koristimo http.ListenAndServe da staramo novi server, dajemo mu dva parametra
	//TCP adresu (:4000) i servemux koji smo upravo stvorili
	//ako http.ListenAndServe baci error, koristimo log.Fatal da izbacimo poruku i izademo iz programa
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
