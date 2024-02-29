package service

import (
	"Personal-Route-Planner/db/dao"
	"Personal-Route-Planner/model"
	"log"
)

type RatingService struct {
	ratingDao *dao.RatingDao
	siteDao   *dao.SiteDao
}

func NewRatingService() *RatingService {
	return &RatingService{ratingDao: dao.NewRatingDao(),
		siteDao: dao.NewSiteDao(),
	}
}

func (r *RatingService) Rating(rate model.Rating) error {
	return r.ratingDao.Rating(rate)
}

func (r *RatingService) CheckRating(name string) (rate string, err error) {
	return r.ratingDao.CheckRating(name)
}

func (r *RatingService) GetRecommand(name string, sid int) (sites []model.Sites, err error) {
	res, err := r.ratingDao.GetRecommand(name, sid)
	if err != nil {
		log.Printf("get recommendation failed...")
		return nil, err
	}
	for _, item := range res {
		site, err := r.siteDao.FindSiteDetails(item + 1)
		if err != nil {
			log.Printf("get sites detailed failed...")
			return nil, err
		}
		sites = append(sites, site)
	}
	return
}
