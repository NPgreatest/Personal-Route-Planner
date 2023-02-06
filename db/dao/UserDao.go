package dao

import "Personal-Route-Planner/model"

type UserDao struct {
	sql []string
}

func NewUserDao() *UserDao {
	return &UserDao{
		sql: []string{
			`SELECT * FROM users WHERE name = ? AND password = ?;`,
			`INSERT INTO comment (cid, name, sid, content, likes, time)VALUES (?, ?, ?, ?, ?, ? );`,
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
