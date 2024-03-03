package dao

type ActivityDao struct {
	sql []string
}

func NewActivityDao() *ActivityDao {
	return &ActivityDao{
		sql: []string{
			`SELECT detail FROM activity WHERE sid=?`,
		}}
}

func (a *ActivityDao) GetActivity(sid int) ([]string, error) {
	var res []string
	err := sqldb.Select(&res, a.sql[0], sid)
	return res, err
}
