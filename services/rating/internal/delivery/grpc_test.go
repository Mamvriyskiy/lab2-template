package delivery

import (
	"context"
	"github.com/golang/mock/gomock"
	"rating/internal"
	models2 "rating/models"
	"rating/proto/rating"
	"testing"
)

func TestGRPCHandler_GetRating(t *testing.T) {
	ctl := gomock.NewController(t)
	name := "test"
	star := 75
	mockUsecase := internal.NewMockRatingUsecase(ctl)
	req := &rating.RatingRequest{
		Name: name,
	}
	uscResponse := models2.UserRatingResponse{
		Stars: uint(star),
	}
	mockUsecase.EXPECT().GetRating(name).Return(uscResponse, models2.OK)
	handler := NewGRPCHandler(mockUsecase)
	res, err := handler.GetRating(context.Background(), req)

	if res.Stars != int32(star) || err != nil {
		t.Error("incorrect result in rating method")
	}
}
