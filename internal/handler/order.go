package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) ViewMain(w http.ResponseWriter, r *http.Request) {

}

// * Get order by uuid
func (h *Handler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	//setting header
	w.Header().Set("Content-Type", "application/json")

	//getting data from url
	OrderUUId := mux.Vars(r)["id"]

	// looking for order by uuid
	order, err := h.s.GetById(OrderUUId)
	if err != nil {
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	//sending request
	msg := Message{
		Order: *order,
	}
	json.NewEncoder(w).Encode(msg)
}
