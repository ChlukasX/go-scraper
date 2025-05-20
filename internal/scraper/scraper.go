package scraper

import (
	"fmt"

	"github.com/ChlukasX/scraper/internal/config"

	"github.com/gocolly/colly"
)

type Scraper struct {
	collector *colly.Collector
	config    *config.Config
}

func (s *Scraper) Start(url string) error {

	s.collector.OnHTML(string(s.config.ScrapeTarget.Name), func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		for _, link := range links {
			fmt.Printf("link: %s\n", link)
		}
	})
	s.collector.Visit(s.config.ScrapeTarget.StartUrl)
	return nil
}

func New(cfg *config.Config) *Scraper {
	c := colly.NewCollector(
		colly.AllowedDomains(cfg.ScrapeTarget.AllowedDomains...),
	)

	return &Scraper{
		collector: c,
		config:    cfg,
	}
}
