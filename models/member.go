package models

type Member struct {
	ID             string       `json:"id"`
	Role           string       `json:"role"`
	Division       string       `json:"division"`
	ActivityCount  int          `json:"activity_count"`
	ActivityHistory []Activity  `json:"activity_history"`
	Achievements   []string    `json:"achievements"`
	FinalProject   *FinalProject `json:"final_project"`
}