package dao

import "Personal-Route-Planner/model"

type UserDao struct {
	sql []string
}

func NewUserDao() *UserDao {
	return &UserDao{
		sql: []string{
			`SELECT * FROM users WHERE id = ? AND password = ?;`,
			`INSERT INTO users (id, name, password, email, avatar, createtime) VALUES (?, ?, ?, ?, ?, ?);`,
			`INSERT INTO comment (cid, uid, sid, content, likes, time)VALUES (?, ?, ?, ?, ?, ? );`,
		},
	}
}

func (u *UserDao) FindUserLogin(id int, password string) (*model.User, error) {
	var user model.User
	err := sqldb.Get(&user, u.sql[0], id, password)
	return &user, err
}

func (u *UserDao) InsertComment(comment model.Comment) error {
	_, err := sqldb.Exec(u.sql[2], comment.Cid, comment.Uid, comment.Sid, comment.Content, comment.Likes, comment.Time)
	return err
}
