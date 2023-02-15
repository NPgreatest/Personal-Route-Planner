package dao

import (
	"Personal-Route-Planner/Algorithm"
	"Personal-Route-Planner/model"
	"encoding/json"
	"errors"
	"fmt"
	matrix "github.com/skelterjohn/go.matrix"
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
		},
	}
}
func (r *RatingDao) GetSitesNumber() (res int, err error) {
	err = sqldb.Get(&res, r.sql[4])
	return
}
func (r *RatingDao) MakeRecommandMatrix() (matrix *matrix.DenseMatrix, err error) {
	js := make([]string, 0)            //查询结果存到string数组
	Rates := make([]model.Rating, 0)   //结构体数组
	smap := make([]map[int]float64, 0) //将大量评价映射到maps中
	tar := make([]float64, 0)          //目标矩阵数组
	rows, err := sqldb.Query(r.sql[3])
	if err != nil {
		return nil, err
	}
	//用Mysql的游标将rates库中所有节点取到js[]数组中，方便后面unmarshall
	for rows.Next() {
		var t string
		err2 := rows.Scan(&t)
		if err2 != nil {
			return nil, err2
		}
		js = append(js, t)
	}
	//将js数组每个string反序列化到model.Rating结构体中
	for _, value := range js {
		var t model.Rating
		err = json.Unmarshal([]byte(string(value)), &t)
		if err != nil {
			return nil, err
		}
		Rates = append(Rates, t)
	}
	//将Rates这个结构体数组中对每个离散景点的评价分散到map中，对之后转换矩阵节省大量的时间复杂度
	for _, value := range Rates {
		t := make(map[int]float64)
		for _, value2 := range value.Rate {
			t[value2.Sid] = value2.Rate
		}
		smap = append(smap, t)
	}
	number, err := r.GetSitesNumber() //获取一共有多少个节点
	if err != nil {
		return nil, err
	}
	for _, value := range smap {
		for i := 1; i <= number; i++ {
			if v, ok := value[i]; ok {
				tar = append(tar, v)
			} else {
				tar = append(tar, 0)
			}
		}
	}
	prefs := Algorithm.MakeRatingMatrix(tar, len(smap), number)
	return prefs, nil
}

func (r *RatingDao) GetRecommand(sid int) (sites []*model.Sites, err error) {
	mat, err := r.MakeRecommandMatrix()
	if err != nil {
		return nil, err
	}
	fmt.Println(mat)
	return
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
