package main

import (
	"be_golang/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler.HandleHello)
	mux.HandleFunc("/nama", handler.HandleNama)
	mux.HandleFunc("/", handler.HandleROot)
	mux.HandleFunc("/data", handler.HandleQuery)
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ini about page"))
	})
	log.Println("startting port on 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
