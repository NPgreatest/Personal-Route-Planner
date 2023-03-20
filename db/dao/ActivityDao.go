package dao

type ActivityDao struct {
	sql []string
}

func NewActivityDao() *ActivityDao {
	return &ActivityDao{
		sql: []string{
			`SELECT name FROM activity WHERE attribute=?`,
		}}
}

func (a *ActivityDao) GetActivity(sid int) ([]string, error) {
	var res []string
	err := sqldb.Select(&res, a.sql[0], sid)
	return res, err
}
