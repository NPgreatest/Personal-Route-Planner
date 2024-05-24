package dao

import (
	"Personal-Route-Planner/Algorithm"
	"Personal-Route-Planner/model"
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
			`SELECT * FROM rating WHERE name=?;`,
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
	// SQL 语句，用于插入和更新操作
	insertSQL := `INSERT INTO route_planner.rating (name, sid, rating) VALUES (?, ?, ?)`
	updateSQL := `UPDATE route_planner.rating SET rating = ? WHERE name = ? AND sid = ?`

	// 开启事务处理，以保证数据一致性
	tx, err := sqldb.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	for _, r := range rate.Rate {
		// 检查是否存在该用户对该景点的评分
		var count int
		err = tx.QueryRow(`SELECT count(*) FROM route_planner.rating WHERE name = ? AND sid = ?`, rate.Name, r.Sid).Scan(&count)
		if err != nil {
			return err
		}

		// 根据查询结果进行插入或更新操作
		if count == 0 {
			_, err = tx.Exec(insertSQL, rate.Name, r.Sid, r.Rate)
		} else {
			_, err = tx.Exec(updateSQL, r.Rate, rate.Name, r.Sid)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RatingDao) CheckRating(name string) (rate []model.RateRes, err error) {
	sql := `SELECT name, sid, rating FROM route_planner.rating WHERE name = ?`
	rows, err := sqldb.Query(sql, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rates []model.RateRes
	for rows.Next() {
		var rate model.RateRes
		err := rows.Scan(&rate.Name, &rate.Sid, &rate.Rate)
		if err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}

	// Check for errors from iterating over rows.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rates, nil
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
