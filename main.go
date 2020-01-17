package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("User making a " + string(r.Method) + " request to " + string(r.URL.Path))

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
	}
}

func data(w http.ResponseWriter, r *http.Request) {
	log.Println("User making a " + string(r.Method) + " request to " + string(r.URL.Path))

	switch r.Method {
	case "GET":
		w.Header().Set("Server", "Go Server")
		w.WriteHeader(200)

		w.Write([]byte("Hello"))

	case "POST":
		log.Println(r.Body)
	}
}

func main() {
	log.Println("Starting server.")

	http.HandleFunc("/", home)
	http.HandleFunc("/data", data)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
