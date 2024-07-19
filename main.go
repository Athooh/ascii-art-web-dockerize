package main

import (
	"fmt"
	"net/http"
	"os"

	handler "web/handlers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}
	http.HandleFunc("/", handler.FormHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/user-guide", handler.HowItWorksHandler)
	http.HandleFunc("/ascii-art", handler.AsciiArtHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
