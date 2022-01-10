package model

type FootballClub struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Tournaments []string `json:"tournaments"`
	Nation      string   `json:"nation"`
	HasStadium  bool     `json:"has_stadium"`
	Description string   `json:"description"`
	Rating      float64  `json:"rating"`
}
