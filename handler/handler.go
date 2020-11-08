package handler

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hai saya sedang belajar golang ww"))
}

func HandleROot(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if url != "/" {
		http.NotFound(w, r)
		return
	}
	// log.Println(url)
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error happenig", http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"nama":    "farhan ammar",
		"content": "saya sedang belajar golang",
	}
	tmpl.Execute(w, data)
}

func HandleNama(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hallo nama saya farhan ammar"))
}

func HandleQuery(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("data ke" + id))
}
