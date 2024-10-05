package utils

import (
    "encoding/json"
    "os"
    "time"

    "github.com/mirula24/bank-api/models"
)

func ReadCustomersFromFile(filename string) ([]models.Customer, error) {
    file, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var customers []models.Customer
    err = json.Unmarshal(file, &customers)
    if err != nil {
        return nil, err
    }

    return customers, nil
}

func WriteCustomersToFile(filename string, customers []models.Customer) error {
    data, err := json.MarshalIndent(customers, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(filename, data, 0644)
}

func LogActivity(activity models.History) error {
    activity.Timestamp = time.Now()

    file, err := os.OpenFile("data/history.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    return encoder.Encode(activity)
}