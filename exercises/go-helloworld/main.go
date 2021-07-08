package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":6111", nil)
}

// https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-18-04
