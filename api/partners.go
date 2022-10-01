package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"api_service/partner"

	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
)

const default_partners_limit = 10 // default limit for DB fetch

type PartnersHandler interface {
	GetAllPartners(w http.ResponseWriter, r *http.Request)
	GetPartner(w http.ResponseWriter, r *http.Request)
	GetFilteredPartners(w http.ResponseWriter, r *http.Request)
}

type partnersHandler struct {
	partnerService partner.Service
}

func NewAPIHandler(
	partnerService partner.Service,
) PartnersHandler {
	return partnersHandler{
		partnerService: partnerService,
	}
}

func (t partnersHandler) GetAllPartners(w http.ResponseWriter, r *http.Request) {
	partners, err := t.partnerService.GetAllPartners()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(echo.HeaderContentType, "application/json")

	json.NewEncoder(w).Encode(partners)
}

func (t partnersHandler) GetPartner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	partnerID, err := strconv.Atoi(params["id"])
	log.Println("Id = ", partnerID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Wrong partner id")
		return
	}

	partner, err := t.partnerService.GetSinglePartner(partnerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Unable to load partner")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(echo.HeaderContentType, "application/json")
	json.NewEncoder(w).Encode(partner)
}

func (t partnersHandler) GetFilteredPartners(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = default_partners_limit
	}

	latitude, err := strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Wrong format of latitude")
		return
	}

	longitude, err := strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Wrong format of longitude")
		return
	}

	materials := r.URL.Query()["materials"]
	partners, err := t.partnerService.GetOrderedPartners(materials, latitude, longitude, limit)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(echo.HeaderContentType, "application/json")
	json.NewEncoder(w).Encode(partners)
}
