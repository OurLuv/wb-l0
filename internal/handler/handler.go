package handler

import (
	"encoding/json"
	"net/http"

	"github.com/OurLuv/l0/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	s service.OrderServcie
}

type Message struct {
	Error string `json:"error,omitempty"`
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/orders", h.ViewMain).Methods("GET")
	r.HandleFunc("/orders/{id}", h.GetOrderById).Methods("GET")

	return r
}

func SendError(w http.ResponseWriter, err string, code int) {
	w.WriteHeader(code)
	response := Message{
		Error: err,
	}
	json.NewEncoder(w).Encode(response)
}

func NewHandler(s service.OrderServcie) *Handler {
	return &Handler{
		s: s,
	}
}
