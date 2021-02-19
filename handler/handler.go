package handler

import (
	"golangbwa/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world, saya sedang belajar golang web"))
}

func UmarHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Umar, Umar sedang belajar guys"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// untuk developer error
	log.Printf(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "i am learning golang web",
		"content": "i am learning golang web with umar",
		"course":  3,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calms", http.StatusInternalServerError)
		return
	}
	//w.Write([]byte("Home"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)
	//nama := r.URL.Query().Get("nama")
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("product page"))
	//fmt.Fprintf(w, "product page : %d dengan nama : %s", idNumb, nama)

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "i am learning golang web",
	// 	"content": idNumb,
	// 	"course":  3,
	// }

	//data := entity.Product{ID: 1, Name: "Mobilio", Price: 9000000, Stock: 4}
	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 9000000, Stock: 12},
		{ID: 2, Name: "heoa", Price: 2000000, Stock: 1},
		{ID: 3, Name: "toyota", Price: 700000, Stock: 2},
		{ID: 4, Name: "forza", Price: 500000, Stock: 7},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calms", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case "GET":
		w.Write([]byte("ini adalah get"))
	case "POST":
		w.Write([]byte("ini adlah post"))
	default:
		http.Error(w, "Error is happening, keep calms", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Erorr is happening, keep calms", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Erorr is happening, keep calms", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error is happening, keep calms", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Erorr is happening, keep calms", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Erorr is happening, keep calms", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Erorr is happening, keep calms", http.StatusInternalServerError)
			return
		}

		//w.Write([]byte(message))
		return

	}

	http.Error(w, "Error is happening, keep calms", http.StatusBadRequest)
}
