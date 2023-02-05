package service

import (
	"Personal-Route-Planner/db/dao"
	"Personal-Route-Planner/model"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{userDao: dao.NewUserDao()}
}

func (u *UserService) CheckUser(id int, password string) (*model.User, error) {
	return u.userDao.FindUserLogin(id, password)
}

func (u *UserService) InsertComment(comment model.Comment) error {
	return u.userDao.InsertComment(comment)
}
