package models

type Language struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation"`
}
