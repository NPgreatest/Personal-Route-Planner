package service

import (
	"Personal-Route-Planner/db/dao"
	"Personal-Route-Planner/model"
)

type HomeService struct {
	homeDao *dao.HomeDao
}

func NewHomeService() *HomeService {
	return &HomeService{homeDao: dao.NewHomeDao()}
}

func (h *HomeService) GetComment() ([]model.Comment, error) {
	return h.homeDao.FindAllComment()
}

func (h *HomeService) RegisterUser(user model.User) error {
	return h.homeDao.RegisterNewUser(user)
}
