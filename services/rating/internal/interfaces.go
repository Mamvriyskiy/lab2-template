package internal

import models2 "rating/models"

type RatingUsecase interface {
	GetRating(name string) (models2.UserRatingResponse, models2.StatusCode)
	UpdateRating(name string, updateNumber int32) models2.StatusCode
}

type RatingRepo interface {
	FetchRating(name string) (models2.UserRatingResponse, models2.StatusCode)
	UpdateRating(name string, updateNumber int32) models2.StatusCode
}
