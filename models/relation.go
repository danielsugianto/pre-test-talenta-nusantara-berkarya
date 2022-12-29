package models

type Relation struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences   []string `json:"influences"`
}
