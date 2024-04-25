package coingecko

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"testtask-oneplus/internal/clients/coingecko/dto"
)

type CoingeckoClient struct {
	//https://api.coingecko.com/api/v3
	url    string
	prices sync.Map
}

func NewCoingeckoClient(url string) *CoingeckoClient {
	return &CoingeckoClient{
		url: url,
	}
}

func (c *CoingeckoClient) getCourses() ([]dto.CourseDTO, error) {
	resp, err := http.Get(c.url + "/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var courses []dto.CourseDTO
	err = json.NewDecoder(resp.Body).Decode(&courses)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CoingeckoClient) UpdateData() error {
	courses, err := c.getCourses()
	if err != nil {
		return err
	}
	for _, course := range courses {
		c.prices.Store(course.ID, course.CurrentPrice)
	}
	return nil
}

func (c *CoingeckoClient) GetPrice(id string) (float64, error) {
	price, ok := c.prices.Load(id)
	if !ok {
		return 0, fmt.Errorf("coingecko: coingecko price not found")
	}
	return price.(float64), nil
}
