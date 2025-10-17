package usecase

import (
	"rating/internal"
	models2 "rating/models"
)

type RatingUsecase struct {
	repo internal.RatingRepo
}

func NewRatingUsecase(repo internal.RatingRepo) *RatingUsecase {
	return &RatingUsecase{repo: repo}
}

func (ru *RatingUsecase) GetRating(name string) (models2.UserRatingResponse, models2.StatusCode) {
	return ru.repo.FetchRating(name)
}

func (ru *RatingUsecase) UpdateRating(name string, updateNumber int32) models2.StatusCode {
	return ru.repo.UpdateRating(name, updateNumber)
}
