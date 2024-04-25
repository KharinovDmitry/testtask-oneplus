package main

import (
	"log"
	"parser/internal/csv"
	"parser/internal/parser"
)

func main() {
	url := "https://hypeauditor.com/top-instagram-all-russia/"
	influencers, err := parser.ParseInfluencer(url)
	if err != nil {
		log.Fatal("Ошибка парсинга: " + err.Error())
	}

	err = csv.SaveInfluencers(influencers, "res")
	if err != nil {
		log.Fatal("Ошибка сохранения: " + err.Error())
	}

}
