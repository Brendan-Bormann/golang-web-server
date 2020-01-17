package main

import (
	"fmt"
	"net/http"
)

func log(status, message string) {
	fmt.Println("[" + status + "] | - " + message)
}

func routeOne() string {
	return "hello"
}

func main() {
	log("Boot", "Starting web server...")

	log("Idle", "---")

	http.HandleFunc("/", routeOne())

	fmt.Scanln()

	log("Stop", "Spinning down application...")
	log("Done", "Press any key to end the application.")
	fmt.Scanln()
}
