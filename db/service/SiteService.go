package service

import (
	"Personal-Route-Planner/db/dao"
	"Personal-Route-Planner/model"
)

type SiteService struct {
	siteDao *dao.SiteDao
}

func NewSiteService() *SiteService {
	return &SiteService{siteDao: dao.NewSiteDao()}
}

func (s *SiteService) GetComment(sid int) ([]model.Comment, error) {
	return s.siteDao.FindAllComment(sid)
}

func (s *SiteService) FindAllSites() (sites []model.Sites, err error) {
	return s.siteDao.FindAllSites()
}

func (s *SiteService) FindSiteDetails(siteid int) (site model.Sites, err error) {
	return s.siteDao.FindSiteDetails(siteid)
}

func (s *SiteService) FindCommentsAvatars(siteid int) ([]string, error) {
	return s.siteDao.FindCommentsAvatars(siteid)
}
