package news

import (
	"fmt"
	"github.com/Gylmynnn/world-news/model"
	"github.com/gocolly/colly"
)

type IndonesiaCountryScraper struct{}

func (s *IndonesiaCountryScraper) Scrape(url, country string, limit, page int) ([]model.News, error) {
	var newsList []model.News
	count := 0
	start := (page - 1) * limit

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(".aimlLatest .wSpec-item", func(e *colly.HTMLElement) {
		title := e.ChildText("a .wSpec-wrap .wSpec-box h4")
		hrefUrl := e.ChildAttr("a", "href")
		desc := "none"
		if count < start {
			count++
			return
		}
		if len(newsList) >= limit {
			return
		}

		newsList = append(newsList, model.News{
			Title:   title,
			Content: desc,
			URL:     e.Request.AbsoluteURL(hrefUrl),
			Source:  url,
			Country: country,
		})
		count++
	})

	err := c.Visit(url)
	return newsList, err
}
