package dao

import (
	"Personal-Route-Planner/model"
	"errors"
	"fmt"
)

type HomeDao struct {
	sql []string
}

func NewHomeDao() *HomeDao {
	return &HomeDao{
		sql: []string{
			`SELECT * FROM comment WHERE sid=?`,
			`INSERT INTO users (name, password, email, avatar, createtime) VALUES ( ?, ?, ?, ?, ?);`,
			`SELECT * FROM sites`,
			`SELECT * FROM sites WHERE sid=?;`,
			`SELECT * FROM price WHERE sid=?;`,
		}}
}

func (h *HomeDao) FindAllComment(sid int) (comments []model.Comment, err error) {
	err = sqldb.Select(&comments, h.sql[0], sid)
	if err != nil {
		return nil, err
	}
	return
}

func (h *HomeDao) RegisterNewUser(user model.User) error {
	_, err := sqldb.Exec(h.sql[1], user.Name, user.Password, user.Email, user.Avatar, user.CreateTime)
	return err
}

func (h *HomeDao) FindAllSites() (sites []model.Sites, err error) {
	err = sqldb.Select(&sites, h.sql[2])
	if err != nil {
		return nil, err
	}
	return
}

func (h *HomeDao) FindSiteDetails(siteid int) (site model.Sites, err error) {
	var prices []model.Price
	var temp []model.Sites
	err = sqldb.Select(&temp, h.sql[3], siteid)
	if err != nil {
		return site, err
	}
	if temp == nil {
		return site, errors.New("site-id wrong")
	}
	site = temp[0]
	fmt.Println(site)
	err = sqldb.Select(&prices, h.sql[4], siteid)
	if err != nil {
		return site, err
	}
	for _, value := range prices {
		site.Prices = append(site.Prices, value)
	}
	return
}
