package model

type Sites struct {
	Sid         int     `json:"sid"`
	Description string  `json:"description"`
	Pic         string  `json:"pic"`
	Website     string  `json:"website"`
	Prices      []Price `json:"prices"`
}
