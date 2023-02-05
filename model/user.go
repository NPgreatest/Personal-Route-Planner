package model

import "time"

type User struct {
	Id         int       `form:"id" json:"id"`
	Name       string    `form:"name" json:"name"`
	Password   string    `form:"password" json:"password"`
	Email      string    `form:"email" json:"email"`
	Avatar     string    `form:"avatar" json:"avatar"`
	CreateTime time.Time `form:"createTime" json:"createTime"`
}
