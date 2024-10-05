package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/mirula24/bank-api/models"
    "github.com/mirula24/bank-api/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var loginRequest struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    err := json.NewDecoder(r.Body).Decode(&loginRequest)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    customers, err := utils.ReadCustomersFromFile("data/customers.json")
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    for _, customer := range customers {
        if customer.Username == loginRequest.Username && customer.Password == loginRequest.Password {
            // Login successful
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
            
            // Log the activity
            utils.LogActivity(models.History{
                Action: "Login",
                Username: customer.Username,
            })
            return
        }
    }

    // Login failed
    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}