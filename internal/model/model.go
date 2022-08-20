package model

type Input struct {
	A float64 `binding:"required" json:"A"`
	B float64 `binding:"required" json:"B"`
}

type Output struct {
	First  float64 `json:"A/B"`
	Second float64 `json:"B/A"`
}
