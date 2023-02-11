package service

import (
	"Personal-Route-Planner/db/dao"
	"Personal-Route-Planner/model"
)

type RatingService struct {
	ratingDao *dao.RatingDao
}

func NewRatingService() *RatingService {
	return &RatingService{ratingDao: dao.NewRatingDao()}
}

func (r *RatingService) Rating(rate model.Rating) error {
	return r.ratingDao.Rating(rate)
}

func (r *RatingService) CheckRating(name string) (rate string, err error) {
	return r.ratingDao.CheckRating(name)
}
