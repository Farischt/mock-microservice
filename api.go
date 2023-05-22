package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/farischt/micro/types"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	service PriceService
	addr    uint
}

func NewJsonApi(service PriceService, addr uint) *Server {
	return &Server{
		service,
		addr,
	}
}

func (s *Server) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
		_ = writeJSON(w, 200, map[string]string{
			"service": "price",
			"status":  "healthy",
		})
	}).Methods("GET")
	r.HandleFunc("/coin/{coin}", makeHttpFunc(s.handlePriceService))

	logrus.WithFields(logrus.Fields{
		"port": s.addr,
	}).Info("JSON Server starting:")

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.addr), r)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to listen and serve:")
	}
}

func (s *Server) handlePriceService(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.getPrice(ctx, w, r)
	default:
		return types.NewApiError(http.StatusMethodNotAllowed, "method_not_allowed")
	}
}

func (s *Server) getPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	coin := vars["coin"]

	price, err := s.service.GetPrice(ctx, coin)

	if err != nil {
		return types.NewApiError(http.StatusBadRequest, err.Error())
	}

	status := http.StatusOK
	priceResponse := types.PriceResponse{
		Coin:  coin,
		Price: price,
	}
	apiResponse := types.NewApiResponse(status, priceResponse, r)
	return writeJSON(w, status, &apiResponse)
}

// Utils
type apiFunc func(
	context.Context,
	http.ResponseWriter,
	*http.Request,
) error

func makeHttpFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(context.Background(), w, r); err != nil {
			// Check errors here
			if e, ok := err.(types.ApiError); ok {
				_ = writeJSON(w, e.Status, e)
				return
			}
			_ = writeJSON(w, http.StatusInternalServerError, types.NewApiError(http.StatusInternalServerError, err.Error()))
		}
	}
}

func writeJSON(w http.ResponseWriter, statusCode int, d any) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(d)
}
