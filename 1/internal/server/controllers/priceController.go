package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"testtask-oneplus/internal/clients/coingecko"
)

type PriceController struct {
	coingeckoClient *coingecko.CoingeckoClient
}

func NewPriceController(coingeckoClient *coingecko.CoingeckoClient) *PriceController {
	return &PriceController{
		coingeckoClient: coingeckoClient,
	}
}

type getPriceResponse struct {
	Price float64 `json:"price"`
}

func (c PriceController) GetPrice(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := r.PathValue("id")
	price, err := c.coingeckoClient.GetPrice(id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response, err := json.Marshal(getPriceResponse{Price: price})
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
