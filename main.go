package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Jenkins and Kubernetes from Go, my name is muhammed essam!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

