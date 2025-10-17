package models

import "time"

type ReturnBookRequest struct {
	Condition BookCondition `json:"condition"`
	Date      time.Time     `json:"date"`
}
