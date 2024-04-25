package server

import (
	"net/http"
	"testtask-oneplus/internal/clients/coingecko"
	"testtask-oneplus/internal/server/controllers"
)

type Server struct {
}

func Run(coingeckoClient *coingecko.CoingeckoClient) error {
	mux := http.NewServeMux()
	priceController := controllers.NewPriceController(coingeckoClient)

	mux.HandleFunc("GET /api/prices/{id}", priceController.GetPrice)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}
	return nil
}
