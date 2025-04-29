package scraper

import (
	"github.com/Gylmynnn/world-news/model"
	"github.com/Gylmynnn/world-news/scraper/news"
)

func GetScraperForCountry(country string) SiteScraper {
	switch country {
	case "japan":
		return &news.JapanCountryScraper{}
	case "indonesian":
		return &news.IndonesiaCountryScraper{}
	case "chinese":
		return &news.ChineseCountryScraper{}
	default:
		return nil
	}
}

type SiteScraper interface {
	Scrape(url string, country string, limit, page int) ([]model.News, error)
}
