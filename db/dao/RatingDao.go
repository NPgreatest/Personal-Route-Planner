package dao

import (
	"Personal-Route-Planner/model"
	"encoding/json"
	"errors"
)

type RatingDao struct {
	sql []string
}

func NewRatingDao() *RatingDao {
	return &RatingDao{
		sql: []string{
			`INSERT INTO rating VALUES (?,?);`,
			`UPDATE rating t SET t.rate=? WHERE t.name=?;`,
			`SELECT rate FROM rating WHERE name=?;`,
		},
	}
}

func (r *RatingDao) Rating(rate model.Rating) error {
	res, _ := json.Marshal(rate)
	sql := `SELECT count(*) FROM rating WHERE name=?`
	var c int
	err := sqldb.Get(&c, sql, rate.Name)
	if err != nil {
		return err
	}
	switch c {
	case 0:
		_, err = sqldb.Exec(r.sql[0], rate.Name, string(res))
	case 1:
		_, err = sqldb.Exec(r.sql[1], string(res), rate.Name)
	default:
		return errors.New("database wrong")
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *RatingDao) CheckRating(name string) (rate string, err error) {
	err = sqldb.Get(&rate, r.sql[2], name)
	if err != nil {
		return
	}
	return
}
