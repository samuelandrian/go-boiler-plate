package users

import (
	"go-boiler-plate/internal/users/delivery"

	"github.com/go-chi/chi"
)

type Routes struct {
	delivery *delivery.Delivery
}

func NewRoutes(delivery *delivery.Delivery) *Routes {
	return &Routes{delivery: delivery}
}

func (routes Routes) RegisterRoutes(r chi.Router) {
	r.Post("/users/greeting", routes.delivery.GreetingHandler)
}
