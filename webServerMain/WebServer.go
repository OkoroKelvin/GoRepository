package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	port := 3000
	port, isStarted := startWebServer(port)
	fmt.Println(port, isStarted)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)
	return port, nil
}
