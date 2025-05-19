package scraper

import (
	"github.com/ChlukasX/scraper/internal/config"
	"github.com/ChlukasX/scraper/internal/storage"

	"github.com/gocolly/colly"
)

type Scraper struct {
	collector *colly.Collector
	config    *config.Config
	storage   storage.DataStore
}

func New(cfg *config.Config, store storage.DataStore) *Scraper {
	c := colly.NewCollector(
		colly.AllowedDomains(),
	)

	return &Scraper{
		collector: c,
		config:    cfg,
		storage:   store,
	}
}
