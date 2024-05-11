package usecase

import (
	"context"
	"go-boiler-plate/internal/users/model"
)

type IService interface {
	Greeting(ctx context.Context, request model.GreetingRequest) (response model.GreetingResponse, err error)
}
