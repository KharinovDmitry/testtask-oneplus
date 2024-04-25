package main

import (
	"log"
	"testtask-oneplus/internal/clients/coingecko"
	"testtask-oneplus/internal/server"
	timer2 "testtask-oneplus/internal/timer"
	"time"
)

func main() {
	client := coingecko.NewCoingeckoClient("https://api.coingecko.com/api/v3")
	timer := timer2.NewTimer(client, 10*time.Minute)
	go timer.Start()

	err := server.Run(client)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
