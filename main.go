package main

import (
    "log"
    "net/http"
)

func main() {
    log.Printf("Initializing...");
    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}


