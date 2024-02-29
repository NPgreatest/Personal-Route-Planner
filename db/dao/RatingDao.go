package dao

import (
	"Personal-Route-Planner/Algorithm"
	"Personal-Route-Planner/model"
	"encoding/json"
	"errors"
	"log"
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
			`SELECT rate FROM rating;`,
			`SELECT count(*) FROM sites;`,
			`SELECT * FROM sites WHERE sid=?;`,
			`SELECT * FROM matrix`,
		},
	}
}

func (r *RatingDao) GetSitesNumber() (res int, err error) {
	err = sqldb.Get(&res, r.sql[4])
	return
}

func (r *RatingDao) GetRecommand(name string, sid int) (sites []int, err error) {
	if len(Algorithm.GlobalVec) == 0 {
		Algorithm.GlobalVec, err = r.GetMatrix()
		if err != nil {
			log.Printf("get matrix failed...")
			return nil, err
		}
	}
	return Algorithm.FindNearestVectors(Algorithm.GlobalVec, sid-1, 5), nil
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

func (r *RatingDao) GetMatrix() (res [][]float64, err error) {
	rows, err := sqldb.Queryx(r.sql[6])
	if err != nil {
		log.Printf("Error querying matrix data: %v", err)
		return nil, err
	}
	defer rows.Close()

	var matrices [][]float64
	for rows.Next() {
		var sid int
		var vec = make([]float64, 10) // 根据表中列的数量调整
		err = rows.Scan(&sid, &vec[0], &vec[1], &vec[2], &vec[3], &vec[4], &vec[5], &vec[6], &vec[7], &vec[8], &vec[9])
		if err != nil {
			log.Printf("Error scanning matrix row: %v", err)
			return nil, err
		}
		matrices = append(matrices, vec)
	}
	return matrices, nil
}
