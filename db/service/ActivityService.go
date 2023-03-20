package service

import "Personal-Route-Planner/db/dao"

type ActivityService struct {
	activityService *dao.ActivityDao
}

func NewActivityService() *ActivityService {
	return &ActivityService{activityService: dao.NewActivityDao()}
}
func (a *ActivityService) GetActivity(sid int) ([]string, error) {
	return a.activityService.GetActivity(sid)
}
