package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/your-actual-username/bank-api/models"
    "github.com/your-actual-username/bank-api/utils"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var paymentRequest struct {
        FromUsername string  `json:"from_username"`
        ToUsername   string  `json:"to_username"`
        Amount       float64 `json:"amount"`
    }

    err := json.NewDecoder(r.Body).Decode(&paymentRequest)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    customers, err := utils.ReadCustomersFromFile("data/customers.json")
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    var fromCustomer, toCustomer *models.Customer
    for i := range customers {
        if customers[i].Username == paymentRequest.FromUsername {
            fromCustomer = &customers[i]
        }
        if customers[i].Username == paymentRequest.ToUsername {
            toCustomer = &customers[i]
        }
    }

    if fromCustomer == nil || toCustomer == nil {
        http.Error(w, "Invalid customer(s)", http.StatusBadRequest)
        return
    }

    // Perform the transfer
    fromCustomer.Balance -= paymentRequest.Amount
    toCustomer.Balance += paymentRequest.Amount

    // Update the customers file
    err = utils.WriteCustomersToFile("data/customers.json", customers)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Log the activity
    utils.LogActivity(models.History{
        Action:   "Payment",
        Username: paymentRequest.FromUsername,
        Details:  fmt.Sprintf("Transferred %.2f to %s", paymentRequest.Amount, paymentRequest.ToUsername),
    })

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Payment successful"})
}