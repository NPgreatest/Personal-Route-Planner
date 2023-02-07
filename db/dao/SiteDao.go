package dao

import (
	"Personal-Route-Planner/model"
	"errors"
)

type SiteDao struct {
	sql []string
}

func NewSiteDao() *SiteDao {
	return &SiteDao{sql: []string{
		`SELECT * FROM comment WHERE sid=?`,
		`SELECT * FROM sites WHERE sid IN(SELECT sid FROM sites_tags WHERE tagid=?);`,
		`SELECT * FROM sites WHERE sid=?;`,
		`SELECT * FROM price WHERE sid=?;`,
		`SELECT avatar FROM users RIGHT JOIN (SELECT name FROM comment WHERE sid=?) a ON a.name=users.name;`,
	}}
}

func (h *SiteDao) FindAllComment(sid int) (comments []model.Comment, err error) {
	err = sqldb.Select(&comments, h.sql[0], sid)
	ava, err := h.FindCommentsAvatars(sid)
	if err != nil {
		return nil, err
	}
	for i, _ := range comments {
		comments[i].Avatar = ava[i]
	}
	//fmt.Println(comments)
	if err != nil {
		return nil, err
	}
	return
}

func (h *SiteDao) FindAllSites(tag int) (sites []model.Sites, err error) {
	err = sqldb.Select(&sites, h.sql[1], tag)
	if err != nil {
		return nil, err
	}
	return
}

func (h *SiteDao) FindSiteDetails(siteid int) (site model.Sites, err error) {
	var prices []model.Price
	var temp []model.Sites
	err = sqldb.Select(&temp, h.sql[2], siteid)
	if err != nil {
		return site, err
	}
	if temp == nil {
		return site, errors.New("site-id wrong")
	}
	site = temp[0]
	//fmt.Println(site)
	err = sqldb.Select(&prices, h.sql[3], siteid)
	if err != nil {
		return site, err
	}
	for _, value := range prices {
		site.Prices = append(site.Prices, value)
	}
	return
}

func (h *SiteDao) FindCommentsAvatars(siteid int) (ava []string, err error) {
	err = sqldb.Select(&ava, h.sql[4], siteid)
	if err != nil {
		return nil, err
	}
	return
}
