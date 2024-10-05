package main

import (
    "log"
    "net/http"

    "github.com/mirula24/bank-api/handlers"
)

func main() {
    http.HandleFunc("/login", handlers.LoginHandler)
    http.HandleFunc("/payment", handlers.PaymentHandler)
    http.HandleFunc("/logout", handlers.LogoutHandler)

    log.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}