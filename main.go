package main

import "fmt"

func log(status, message string) {
	fmt.Println("[" + status + "] | - " + message)
}

func main() {
	log("Boot", "Starting web server...")

	log("Idle", "...")

	fmt.Scanln()

	log("Stop", "Spinning down application...")
	log("Done", "Press any key to end the application.")
	fmt.Scanln()
}
