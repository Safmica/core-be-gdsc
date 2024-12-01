package models

type FinalProject struct {
	ID          string `json:"id"`
	DivisionID string `json:"division_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
