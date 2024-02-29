package model

type Rating struct {
	Name string  `json:"name" binding:"required"`
	Rate []Rates `json:"rates" binding:"dive"`
}

type Rates struct {
	Sid  int     `json:"sid" binding:"required"`
	Rate float64 `json:"rate" binding:"required"`
}
type MatrixRow struct {
	SID int       `db:"sid"`
	Vec []float64 // 这个不直接映射数据库字段，需要后续处理
}
