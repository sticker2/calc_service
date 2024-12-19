package main

import (
    "log"
    "net/http"
    "calc_service/internal/application"
)

func main() {
    http.HandleFunc("/api/v1/calculate", application.HelloHandler)
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
