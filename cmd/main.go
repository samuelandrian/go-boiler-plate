package main

import (
	"fmt"
	"go-boiler-plate/internal/users"
	usersDelivery "go-boiler-plate/internal/users/delivery"
	"go-boiler-plate/internal/users/usecase"
	"go-boiler-plate/pkg/config"
	"go-boiler-plate/pkg/timehelper"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	config.Config.SetEnvironment()
	config.Config.ReadConfigFromFile()
	fmt.Println("listening on port :", config.Config.App.Port)
	fmt.Println(config.Config.App.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", config.Config.App.Port), newRoutes())
}

func newRoutes() http.Handler {
	route := chi.NewRouter()
	injectRoutesUsers(route)
	return http.TimeoutHandler(route, timehelper.TIMEOUT_DURATION, "{\"Message\": \"Service Unavailable\"}")
}

func injectRoutesUsers(r chi.Router) {
	service := usecase.NewServiceImplementation(usecase.ServiceOption{Config: config.Config})
	delivery := usersDelivery.NewDelivery(service)
	route := users.NewRoutes(delivery)
	route.RegisterRoutes(r)
}
