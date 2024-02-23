package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func (h *Handler) ViewMain(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		errStr := fmt.Sprintf("can't load a view 01: %s", err.Error())
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	if err = t.Execute(w, nil); err != nil {
		errStr := fmt.Sprintf("can't load a view 02: %s", err.Error())
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
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
	json.NewEncoder(w).Encode(order)
}
