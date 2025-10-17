package utils

import models2 "reservation/models"

func Validate(req models2.TakeBookRequest) *models2.ValidationErrorResponse {
	return &models2.ValidationErrorResponse{}
}
