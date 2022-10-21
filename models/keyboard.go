package models

type Keyboard struct {
	Id           string  `json:"id"`
	Model        string  `json:"model"`
	Manufacturer string  `json:"manufacturer"`
	Price        float64 `json:"price"`
}
