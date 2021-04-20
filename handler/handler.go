package handler

import (
	"crudgolang/entity"
	"html/template"
	"log"
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { //Jika URL tidak ada maka 404 Not Found
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error 505", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title":   "Golang Website",
		"content": "Bahasa Pemprograman Golang",
	}

	tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error 505", http.StatusInternalServerError)
		return
	}
}

func HiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Word"))
}

func SettingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Menu Setting"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	//	id := r.URL.Query().Get("id")

	//idNumb, err := strconv.Atoi(id)
	// if err != nil || idNumb < 1 { //Filtering Input Pengguna
	// 	http.NotFound(w, r)
	// 	return
	// }
	// data := map[string]interface{}{
	// 	"content": idNumb,
	// }

	data := entity.Product{ID: 1, Name: "Mobilio", Price: 2200000, Stock: 3}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error 505", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error 505", http.StatusInternalServerError)
		return
	}
}
