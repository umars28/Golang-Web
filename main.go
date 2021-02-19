package main

import (
	"golangbwa/handler"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()

	// closure
	// aboutHandler := func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("about page"))
	// }

	//mux.HandleFunc("/about", aboutHandler)
	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("profil page"))
	// })
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/umar", handler.UmarHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.Process)

	// load static file
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//localhost:8080/static/style.css
	///static/style.css

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
