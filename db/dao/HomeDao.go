package dao

import (
	"Personal-Route-Planner/model"
)

type HomeDao struct {
	sql []string
}

func NewHomeDao() *HomeDao {
	return &HomeDao{
		sql: []string{
			`INSERT INTO users (name, password, email, avatar, createtime) VALUES ( ?, ?, ?, ?, ?);`,
		}}
}
func (h *HomeDao) RegisterNewUser(user model.User) error {
	_, err := sqldb.Exec(h.sql[0], user.Name, user.Password, user.Email, user.Avatar, user.CreateTime)
	return err
}
