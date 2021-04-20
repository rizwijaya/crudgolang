package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //Create Routing

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", hiHandler)
	mux.HandleFunc("/setting", settingHandler)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux) //Starting on Port 80
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { //Jika URL tidak ada maka 404 Not Found
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to Home"))
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Word"))
}

func settingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Menu Setting"))
}
