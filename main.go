package main

import (
	"log"
	"net/http"
	"time"
)

func home(res http.ResponseWriter, req *http.Request) {
	log.Println(string(req.Method) + " request to " + string(req.URL.Path))

	switch req.Method {
	case "GET":
		http.ServeFile(res, req, "index.html")
	}
}

func data(res http.ResponseWriter, req *http.Request) {
	log.Println(string(req.Method) + " request to " + string(req.URL.Path))

	switch req.Method {
	case "GET":
		res.Header().Set("Server", "Go Server")
		res.WriteHeader(200)

		res.Write([]byte("Hello"))

	case "POST":
		err := req.ParseForm()

		if err != nil {
			log.Panic(err)
		}

		data := req.Form
		log.Println("POST Data: ", data)
	}
}

func doWork(message chan string) {

	message <- "Starting Job"

	for i := 0; i < 5; i++ {
		log.Printf("Work step %v", i+1)
		time.Sleep(1 * time.Second)
	}

	message <- "Job's done."
}

func work(res http.ResponseWriter, req *http.Request) {
	log.Println(string(req.Method) + " request to " + string(req.URL.Path))

	res.Header().Set("Server", "Go Server")
	res.WriteHeader(200)

	message := make(chan string)

	go doWork(message)

	log.Println(<-message)

	log.Println(<-message)

	res.Write([]byte("Work Complete"))
}

func main() {
	log.Println("Starting server.")

	http.HandleFunc("/", home)
	http.HandleFunc("/data", data)
	http.HandleFunc("/work", work)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
