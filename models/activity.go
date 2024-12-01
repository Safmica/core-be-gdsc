package models

type Activity struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        string    `json:"date"`
}