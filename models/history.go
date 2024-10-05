package models

import "time"

type History struct {
    Timestamp time.Time `json:"timestamp"`
    Action    string    `json:"action"`
    Username  string    `json:"username"`
    Details   string    `json:"details,omitempty"`
}