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
	base_url := flag.String("url", "https://en.wikipedia.org/wiki/Web_scraping", "The base URL to scrape from")
	cfg_path := flag.String("config", "./config.yaml", "path to config file")
	csv_headers := flag.String("csv-headers", "", "path to csv with headers that will be scraped")
	colly := flag.Bool("colly", false, "Boolean for if you want to use colly (default false)")

	flag.Parse()

	cfg, err := config.LoadConfig(cfg_path)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)

	dataStore, err := storage.New("sqlite")
	if err != nil {
		log.Fatalf("Error creating data store: %v", err)
	}
	defer dataStore.Close()

	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if *colly == true {
		log.Println("using colly...")
		use_colly(*base_url)
		colly_tables(*writer)
		return
	}

	log.Print(*base_url)

	headers := [...]string{
		"Dataset Name", "Description",
	}

	if len(*csv_headers) == 0 {
		log.Printf("Using default headers: %s", headers)
	} else {
		_, err := read_csv(*csv_headers)
		if err != nil {
			log.Fatal(err)
		}
	}

	resp, err := fetch_site(*base_url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(body))

	if rootNode, err := html.Parse(resp.Body); err == nil {
		for _, link := range findLinks(rootNode) {
			log.Printf("link: %s\n", link)
		}
	}
	defer resp.Body.Close()
}

func colly_tables(writer csv.Writer) {
	c := colly.NewCollector()
	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			writer.Write([]string{
				el.ChildText("td:nth-child(1)"),
				el.ChildText("td:nth-child(2)"),
				el.ChildText("td:nth-child(3)"),
			})
		})
		fmt.Println("Scrapping Complete")
	})
	c.Visit("https://www.w3schools.com/html/html_tables.asp")
}

func use_colly(url string) {

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		for _, link := range links {
			log.Printf("link: %s\n", link)
		}
	})
	c.Visit(url)
}

func fetch_site(url string) (*http.Response, error) {
	if resp, err := http.Get(url); err == nil {
		return resp, nil
	} else {
		return nil, err
	}
}

func findLinks(n *html.Node) (links []string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, findLinks(c)...)
	}
	return links
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
