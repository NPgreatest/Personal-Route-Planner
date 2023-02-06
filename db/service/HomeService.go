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

func (h *HomeService) GetComment(sid int) ([]model.Comment, error) {
	return h.homeDao.FindAllComment(sid)
}

func (h *HomeService) RegisterUser(user model.User) error {
	return h.homeDao.RegisterNewUser(user)
}

func (h *HomeService) FindAllSites() (sites []model.Sites, err error) {
	return h.homeDao.FindAllSites()
}

func (h *HomeService) FindSiteDetails(siteid int) (site model.Sites, err error) {
	return h.homeDao.FindSiteDetails(siteid)
}
