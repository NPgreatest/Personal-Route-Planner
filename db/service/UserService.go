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

func (u *UserService) CheckUser(name string, password string) (*model.User, error) {
	return u.userDao.FindUserLogin(name, password)
}

func (u *UserService) InsertComment(comment model.Comment) error {
	return u.userDao.InsertComment(comment)
}

func (u *UserService) FindName(name string) (bool, error) {
	return u.userDao.FindName(name)
}
