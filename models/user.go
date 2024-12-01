package models

type User struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Email         string      `json:"email"`
	University    string      `json:"university"`
	Major         string      `json:"major"`
	Year          string      `json:"year"`
	NIM           string      `json:"nim"`
	CurrentBatch  int         `json:"current_batch"`
	BatchHistories []Batch    `json:"batch_history"`
}