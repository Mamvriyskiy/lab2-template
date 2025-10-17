package delivery

import (
	"context"
	"errors"
	"fmt"
	"rating/internal"
	"rating/models"
	"rating/proto/rating"
)

type GRPCHandler struct {
	usecase internal.RatingUsecase
}

func NewGRPCHandler(usecase internal.RatingUsecase) *GRPCHandler {
	return &GRPCHandler{usecase: usecase}
}

func (h *GRPCHandler) GetRating(ctx context.Context, req *rating.RatingRequest) (*rating.RatingResponse, error) {
	name := req.Name
	rate, status := h.usecase.GetRating(name)
	if status != models.OK {
		return nil, errors.New(fmt.Sprintf("%d", status))
	}

	return &rating.RatingResponse{Stars: int32(rate.Stars)}, nil
}

func (h *GRPCHandler) RatingUpdate(ctx context.Context, req *rating.RatingUpdateRequest) (*rating.RatingUpdateResponse, error) {
	name := req.Name
	number := req.Add
	st := h.usecase.UpdateRating(name, number)
	if st != models.OK {
		return nil, errors.New(fmt.Sprintf("%d", st))
	}
	return &rating.RatingUpdateResponse{Ok: true}, nil
}
