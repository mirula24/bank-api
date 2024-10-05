package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/mirula24/bank-api/models"
    "github.com/mirula24/bank-api/utils"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var logoutRequest struct {
        Username string `json:"username"`
    }

    err := json.NewDecoder(r.Body).Decode(&logoutRequest)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Log the activity
    utils.LogActivity(models.History{
        Action:   "Logout",
        Username: logoutRequest.Username,
    })

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}