package model

type Rating struct {
	Name string
	Rate []Rates `json:"rates" binding:"dive"`
}

type Rates struct {
	Sid  int     `json:"sid" binding:"required"`
	Rate float64 `json:"rate" binding:"required"`
}
