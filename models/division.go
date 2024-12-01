package models

type Division struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	MemberIDs      []string    `json:"members"`
	ActivityCount  int         `json:"activity_count"`
	ActivityHistory []string   `json:"activity_history"`
	Documentations []string    `json:"documentations"`
	FinalProjects  []FinalProject `json:"final_projects"`
}