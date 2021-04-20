package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { //Jika URL tidak ada maka 404 Not Found
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to Home"))
}

func HiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Word"))
}

func SettingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Menu Setting"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 { //Filtering Input Pengguna
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Product id Number: %d", idNumb)
}
