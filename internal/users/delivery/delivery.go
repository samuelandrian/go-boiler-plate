package delivery

import (
	"encoding/json"
	"go-boiler-plate/internal/users/model"
	"go-boiler-plate/internal/users/usecase"
	"go-boiler-plate/pkg/responsehelper"
	"net/http"
)

type Delivery struct {
	svc usecase.IService
}

func NewDelivery(svc usecase.IService) *Delivery {
	return &Delivery{
		svc: svc,
	}
}

func (d Delivery) GreetingHandler(w http.ResponseWriter, r *http.Request) {
	var request model.GreetingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responsehelper.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	response, err := d.svc.Greeting(r.Context(), request)
	if err != nil {
		responsehelper.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responsehelper.SendResponse(w, http.StatusOK, response)
}
