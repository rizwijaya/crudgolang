package main

import (
	"crudgolang/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //Create Routing

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HiHandler)
	mux.HandleFunc("/setting", handler.SettingHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux) //Starting on Port 80
	log.Fatal(err)
}
