package dao

import (
	"Personal-Route-Planner/model"
	"fmt"
)

type UserDao struct {
	sql []string
}

func NewUserDao() *UserDao {
	return &UserDao{
		sql: []string{
			`SELECT * FROM users WHERE name = ? AND password = ?;`,
			`INSERT INTO comment (cid, name, sid, content, likes, time)VALUES (?, ?, ?, ?, ?, ? );`,
			`SELECT count(*) FROM users WHERE users.name=?`,
		},
	}
}

func (u *UserDao) FindUserLogin(name string, password string) (*model.User, error) {
	var user model.User
	err := sqldb.Get(&user, u.sql[0], name, password)
	return &user, err
}

func (u *UserDao) InsertComment(comment model.Comment) error {
	_, err := sqldb.Exec(u.sql[1], comment.Cid, comment.Name, comment.Sid, comment.Content, comment.Likes, comment.Time)
	return err
}

func (u *UserDao) FindName(name string) (bool, error) {
	var res int64
	err := sqldb.Get(&res, u.sql[2], name)
	if err != nil {
		return false, err
	}
	fmt.Println(res)
	return res == 0, err
}
