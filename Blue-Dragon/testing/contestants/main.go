package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Serve static files from the "public" directory
    fs := http.FileServer(http.Dir("public"))
    http.Handle("/", fs)

    // Serve the "index.html" file as the root URL
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "public/index.html")
    })

    // Start the server
    fmt.Println("Server listening on port 3000")
    http.ListenAndServe(":3000", nil)
}
