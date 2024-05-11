package model

import (
	"go-boiler-plate/pkg/errorhelper"
)

type GreetingRequest struct {
	Name  string `json:"name" example:"samuel"`
	Date  int    `json:"date" example:"12"`
	Month int    `json:"month" example:"12"`
	Year  int    `json:"year" example:"1993"`
}

func (request GreetingRequest) Validate() (err error) {
	if request.Name == "" {
		return errorhelper.ErrorInvalidRequest
	}
	if request.Date < 1 && request.Date > 31 {
		return errorhelper.ErrorInvalidRequest
	}
	if request.Month < 1 && request.Month > 12 {
		return errorhelper.ErrorInvalidRequest
	}
	if request.Year < 1900 {
		return errorhelper.ErrorInvalidRequest
	}
	return
}

type GreetingResponse struct {
	Name      string `json:"name" example:"S10H"`
	AgeYears  int    `json:"ageYears" example:"12"`
	AgeMonths int    `json:"ageMonths" example:"10"`
	AgeDays   int    `json:"ageDays" example:"10"`
	Zodiac    string `json:"zodiac" example:"capricorn"`
	Text      string `json:"text" example:"capricorn"`
	Env       string `json:"env"`
}
