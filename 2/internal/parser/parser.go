package parser

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"parser/internal/domain"
	"strings"
)

func ParseInfluencer(url string) ([]domain.Influencer, error) {
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		return nil, err
	}
	defer service.Stop()

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	caps.AddChrome(chrome.Capabilities{Args: []string{}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return nil, err
	}

	err = driver.MaximizeWindow("")
	if err != nil {
		return nil, err
	}

	err = driver.Get(url)
	if err != nil {
		return nil, err
	}

	els, err := driver.FindElements(selenium.ByCSSSelector, ".row__top")
	if err != nil {
		return nil, err
	}

	influencers := make([]domain.Influencer, len(els))
	for i := 0; i < len(els); i++ {
		rank, err := getTextEl(els[i], ".rank")
		if err != nil {
			return nil, err
		}
		subscribes, err := getTextEl(els[i], ".subscribers")
		if err != nil {
			return nil, err
		}
		audience, err := getTextEl(els[i], ".audience")
		if err != nil {
			return nil, err
		}
		authentic, err := getTextEl(els[i], ".authentic")
		if err != nil {
			return nil, err
		}
		engagement, err := getTextEl(els[i], ".engagement")
		if err != nil {
			return nil, err
		}

		contributorEl, err := els[i].FindElement(selenium.ByCSSSelector, ".contributor")
		if err != nil {
			return nil, err
		}
		nick, err := getTextEl(contributorEl, ".contributor__name")
		if err != nil {
			return nil, err
		}
		nick = strings.SplitN(nick, "\n", 2)[0]

		name, err := getTextEl(contributorEl, ".contributor__title")
		if err != nil {
			return nil, err
		}

		categoriesEl, err := els[i].FindElement(selenium.ByCSSSelector, ".category")
		if err != nil {
			return nil, err
		}
		categoryEls, err := categoriesEl.FindElements(selenium.ByCSSSelector, ".tag")
		if err != nil {
			return nil, err
		}
		categories := make([]string, len(categoryEls))
		for j, categoryEl := range categoryEls {
			text, err := categoryEl.Text()
			if err != nil {
				return nil, err
			}
			categories[j] = text
		}

		influencers[i] = domain.Influencer{
			Rank:       rank,
			Subscribes: subscribes,
			Audience:   audience,
			Authentic:  authentic,
			Engagement: engagement,
			Nick:       nick,
			Name:       name,
			Categories: categories,
		}
	}
	return influencers, nil
}

func getTextEl(el selenium.WebElement, selector string) (string, error) {
	subEl, err := el.FindElement(selenium.ByCSSSelector, selector)
	if err != nil {
		return "", err
	}
	text, err := subEl.Text()
	if err != nil {
		return "", err
	}
	return text, nil
}
