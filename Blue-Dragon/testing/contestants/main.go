package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Serve static files from the "public" directory
    fs := http.FileServer(http.Dir("public"))
    http.Handle("/", fs)

    // Start the server
    fmt.Println("Server listening on port 3000")
    http.ListenAndServe(":3000", nil)
}
