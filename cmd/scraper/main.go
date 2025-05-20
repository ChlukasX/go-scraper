package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ChlukasX/scraper/internal/config"
	"github.com/ChlukasX/scraper/internal/scraper"
)

func main() {
	cfg_path := flag.String("config", "./config.yaml", "path to config file")

	flag.Parse()

	cfg, err := config.New(*cfg_path)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Print(cfg.ScrapeTarget)

	scraper := scraper.New(cfg)

	err = scraper.Start(cfg.ScrapeTarget.StartUrl)
	if err != nil {
		log.Fatalf("Scraping error: %v", err)
	}

	log.Println("Scraping finished!")
}
