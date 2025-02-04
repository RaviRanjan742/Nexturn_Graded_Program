package models

import "time"

type Blog struct {
    ID        int64     `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Author    string    `json:"author"`
    Timestamp time.Time `json:"timestamp"`
}