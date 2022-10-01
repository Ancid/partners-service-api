package api

import (
	"github.com/gorilla/mux"

	"api_service/database"
	"api_service/partner"

	"gorm.io/gorm"
)

type Handler struct {
	api PartnersHandler
}

func New(db *gorm.DB) *Handler {
	repository := database.NewRepository(db)
	partnerService := partner.NewService(repository)
	api := NewAPIHandler(partnerService)
	h := Handler{api}

	return &h
}

func CreateRouter(h *Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/partners", h.api.GetFilteredPartners).Methods("GET")
	r.HandleFunc("/partners/{id:[0-9]+}", h.api.GetPartner).Methods("GET")
	r.HandleFunc("/partners/list", h.api.GetAllPartners).Methods("GET")

	return r
}
