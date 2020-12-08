package main

import (
	"bytes"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Title         string `yaml:"title"`
	CSSInclude    string `yaml:"css"`
	JSInclude     string `yaml:"js"`
	HeaderContent string `yaml:"header"`
	FooterContent string `yaml:"footer"`
	ImagesWidth   string `yaml:"images-width"`
	ImagesHeight  string `yaml:"images-height"`

	Leader  Slide   `yaml:"leader"`
	Slides  []Slide `yaml:"slides"`
	Trailer Slide   `yaml:"trailer"`
}

func loadConfig() *Config {
	data, _ := ioutil.ReadFile(*oConfig)
	dec := yaml.NewDecoder(bytes.NewReader(data))
	config := new(Config)
	_ = dec.Decode(config)
	return config
}
