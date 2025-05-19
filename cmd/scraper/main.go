package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ChlukasX/scraper/internal/config"
	"github.com/ChlukasX/scraper/internal/storage"

	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)

func main() {
	cfg_path := flag.String("config", "./config.yaml", "path to config file")

	flag.Parse()

	cfg, err := config.LoadConfig(cfg_path)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)

	dataStore, err := storage.New("sqlite")
	if err != nil {
		log.Fatalf("Error creating data store: %v", err)
	}
	defer dataStore.Close()

	scraper := Scraper.New(cfg, dataStore)

	err = scraper.Start(cfg.StartUrl)
	if err != nil {
		log.Fatalf("Scraping error: %v", err)
	}

	log.Println("Scraping finished!")
}
