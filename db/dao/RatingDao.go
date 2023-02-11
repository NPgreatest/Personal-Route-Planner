package dao

import (
	"Personal-Route-Planner/model"
	"encoding/json"
	"fmt"
)

type RatingDao struct {
	sql []string
}

func NewRatingDao() *RatingDao {
	return &RatingDao{
		sql: []string{
			`INSERT INTO rating VALUES (?,?);`,
			`SELECT rate FROM rating WHERE name=?;`,
		},
	}
}

func (r *RatingDao) Rating(rate model.Rating) error {
	res, _ := json.Marshal(rate)
	fmt.Println(string(res))
	_, err := sqldb.Exec(r.sql[0], rate.Name, Convet(rate))
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (r *RatingDao) CheckRating(name string) (rate model.Rating, err error) {
	err = sqldb.Get(&rate, r.sql[1], name)
	fmt.Println(rate)
	if err != nil {
		return
	}
	return
}

func Convet(rate model.Rating) string {

	return ""
}
