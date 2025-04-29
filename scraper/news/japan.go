package news

import (
	"fmt"

	"github.com/Gylmynnn/world-news/model"
	"github.com/gocolly/colly"
)

type JapanCountryScraper struct{}

func (s *JapanCountryScraper) Scrape(url, country string, limit, page int) ([]model.News, error) {
	var newsList []model.News
	count := 0
	start := (page - 1) * limit

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(".article-details", func(e *colly.HTMLElement) {
		title := e.ChildText(".article-title a")
		hrefUrl := e.ChildAttr(".article-title a", "href")
		desc := e.ChildText(".article-body a")
		if desc == "" {
			desc = "none"
		}
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
