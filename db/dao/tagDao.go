package dao

import (
	"Personal-Route-Planner/model"
	"fmt"
)

type TagDao struct {
	sql []string
}

func NewTagDao() *TagDao {
	return &TagDao{sql: []string{
		`SELECT * FROM tags`,
		`SELECT distinct sites.sid,sname,description,pic,website FROM sites,sites_tags WHERE sites.sid=sites_tags.sid AND sites_tags.tagid=?`,
		`SELECT name FROM tags WHERE tags.tagid=?`,
	}}
}

func (h *TagDao) FindAllTags() (tags []model.TagList, err error) {
	err = sqldb.Select(&tags, h.sql[0])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return
}

func (h *TagDao) FindSitesByTags(tagid int) (sites []model.Sites, err error) {
	err = sqldb.Select(&sites, h.sql[1], tagid)
	if err != nil {
		return nil, err
	}
	return
}

func (h *TagDao) FindTagName(tagid int) (res string, err error) {
	err = sqldb.Select(&res, h.sql[2], tagid)
	return res, err
}
