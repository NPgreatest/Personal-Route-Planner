package service

import (
	"Personal-Route-Planner/db/dao"
	"Personal-Route-Planner/model"
)

type TagService struct {
	tagDao *dao.TagDao
}

func NewTagService() *TagService {
	return &TagService{
		tagDao: dao.NewTagDao(),
	}
}

func (t *TagService) FindAllTags() (tags []model.TagList, err error) {
	return t.tagDao.FindAllTags()
}

func (t *TagService) FindSitesByTags(tagid int) (sites []model.Sites, err error) {
	return t.tagDao.FindSitesByTags(tagid)
}
