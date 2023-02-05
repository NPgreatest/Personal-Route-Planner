package dao

import "Personal-Route-Planner/model"

type HomeDao struct {
	sql []string
}

func NewHomeDao() *HomeDao {
	return &HomeDao{
		sql: []string{
			`SELECT * FROM comment`,
			`INSERT INTO users (id, name, password, email, avatar, createtime) VALUES (?, ?, ?, ?, ?, ?);`,
			`SELECT * FROM sites`,
		}}
}

func (h *HomeDao) FindAllComment() (comments []model.Comment, err error) {
	err = sqldb.Select(&comments, h.sql[0])
	if err != nil {
		return nil, err
	}
	return
}

func (h *HomeDao) RegisterNewUser(user model.User) error {
	_, err := sqldb.Exec(h.sql[1], user.Id, user.Name, user.Password, user.Email, user.Avatar, user.CreateTime)
	return err
}
