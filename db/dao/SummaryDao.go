package dao

import (
	"math/rand"
	"time"
)

type SummaryDao struct {
	sql []string
}

func NewSummaryDao() *SummaryDao {
	return &SummaryDao{sql: []string{
		`SELECT summary FROM summary WHERE tag=?`,
	}}
}

func (s *SummaryDao) GetSummary(tagid int) (string, error) {
	rows, err := sqldb.Query(s.sql[0], tagid)
	if err != nil {
		return "", err
	}
	arr := make([]string, 0)
	for rows.Next() {
		var t string
		err2 := rows.Scan(&t)
		if err2 != nil {
			return "", err2
		}
		arr = append(arr, t)
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(arr))
	return arr[randomIndex], nil
}
