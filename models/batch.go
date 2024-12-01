package models

type Batch struct {
	ID          string      `json:"id"`
	Year        int         `json:"year"`
	MemberIDs   []string    `json:"members"`
	DivisionIDs []string    `json:"divisions"`
	Activities  []Activity  `json:"activities"`
}