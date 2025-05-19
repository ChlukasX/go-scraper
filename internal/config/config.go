package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ScrapeTarget struct {
	Name           string   `yaml:"name"`
	StartUrl       string   `yaml:"start_url"`
	AllowedDomains []string `yaml:"allowed_domains"`
	ItemSelector   string   `yaml:"item_selector"`
}

type Config struct {
	ScrapeTarget []ScrapeTarget `yaml:"scrape_targets"`
}

func LoadConfig(path string) (*Config, error) {
	cfgf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(cfgf, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
