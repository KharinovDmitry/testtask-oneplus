package timer

import (
	"log"
	"testtask-oneplus/internal/clients/coingecko"
	"time"
)

type Timer struct {
	coingeckoClient *coingecko.CoingeckoClient
	interval        time.Duration
}

func NewTimer(coingeckoClient *coingecko.CoingeckoClient, interval time.Duration) *Timer {
	return &Timer{
		coingeckoClient: coingeckoClient,
		interval:        interval,
	}
}

func (t *Timer) Start() {
	for {
		err := t.coingeckoClient.UpdateData()
		if err != nil {
			log.Printf("Error updating coingecko data: %v", err)
		} else {
			log.Printf("Updated coingecko data: %v", time.Now())
		}

		time.Sleep(t.interval)
	}
}
