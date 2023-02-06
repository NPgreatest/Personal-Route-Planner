package model

type Sites struct {
	Sid         int     `json:"sid"`
	Sname       string  `json:"sname"`
	Description string  `json:"description"`
	Pic         string  `json:"pic"`
	Website     string  `json:"website"`
	Prices      []Price `json:"prices"`
}
