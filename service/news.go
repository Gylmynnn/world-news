package service

import (
	"errors"

	"github.com/Gylmynnn/world-news/model"
	"github.com/Gylmynnn/world-news/scraper"
)

type countrySource struct {
}

func getUrlForCountry(country string) string {
	switch country {
	case "japan":
		return "https://www.japantimes.co.jp/"
	case "indonesian":
		return "https://www.kompas.com/"
	case "chinese":
		return "https://www.chinadaily.com.cn/"
	default:
		return ""
	}
}

// var countrySources = map[string]string{
// 	"indonesia": "https://www.cnnindonesia.com/",
// 	"japan":     "https://www.japantimes.co.jp/",
// 	"chinese":   "https://www.chinadaily.com.cn/",
// }

func ScrapNewsByCountry(country string, limit, page int) ([]model.News, error) {
	scraperImpl := scraper.GetScraperForCountry(country)
	if scraperImpl == nil {
		return nil, errors.New("country not supported")
	}

	url := getUrlForCountry(country)
	if url == "" {
		return nil, errors.New("URL not found for country")
	}

	return scraperImpl.Scrape(url, country, limit, page)
}
