package model

type Price struct {
	Sid   int     `json:"sid"`
	Aimed string  `json:"aimed"`
	Price float64 `json:"price"`
}
