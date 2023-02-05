package model

import "time"

type Comment struct {
	Cid     int64     `form:"cid" json:"cid"`
	Uid     int       `form:"uid"json:"uid"`
	Sid     int       `form:"sid"json:"sid"`
	Content string    `form:"content"json:"content"`
	Likes   int       `form:"likes"json:"likes"`
	Time    time.Time `form:"time"json:"time"`
}
