package config

import (
	"fmt"
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
	ScrapeTarget ScrapeTarget `yaml:"scrape_targets"`
}

func New(path string) (*Config, error) {
	cfg := &Config{}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d := yaml.NewDecoder(f)

	if err := d.Decode(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !s.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}
	return nil
}
