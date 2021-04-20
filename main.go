package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux() //Create Routing

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	}

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", hiHandler)
	mux.HandleFunc("/setting", settingHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halaman Profile"))
	})
	mux.HandleFunc("/product", productHandler)

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

func productHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 { //Filtering Input Pengguna
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Product id Number: %d", idNumb)
}
