package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	base_url := flag.String("url", "https://archive.ics.uci.edu/datasets", "The base URL to scrape from")
	csv := flag.String("csv-headers", "", "path to csv with headers that will be scraped")

	flag.Parse()

	log.Print(*base_url)

	headers := [...]string{
		"Dataset Name", "Description",
	}

	if len(*csv) == 0 {
		log.Printf("Using default headers: %s", headers)
	} else {
		_, err := read_csv(*csv)
		if err != nil {
			log.Fatal(err)
		}
	}

	body, err := fetch_site(*base_url)
	if err != nil {
		log.Fatal(err)
	}
	println(string(body))
}

func fetch_site(url string) ([]byte, error) {
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {
			return body, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func read_csv(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.Read()
	if err != nil {
		return nil, err
	}

	return records, nil
}
