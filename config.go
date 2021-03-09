package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Title         string        `yaml:"title"`
	CSSInclude    string        `yaml:"css"`
	JSInclude     string        `yaml:"js"`
	HeaderContent template.HTML `yaml:"header"`
	FooterContent template.HTML `yaml:"footer"`
	ImagesWidth   string        `yaml:"images-width"`
	ImagesHeight  string        `yaml:"images-height"`

	PauseBeforeAudio time.Duration `yaml:"pause-before-audio"`
	PauseAfterAudio  time.Duration `yaml:"pause-after-audio"`

	Leader  Slide   `yaml:"leader"`
	Slides  []Slide `yaml:"slides"`
	Trailer Slide   `yaml:"trailer"`
}

func loadConfig() *Config {
	data, err := ioutil.ReadFile(*oConfig)
	check(err)
	dec := yaml.NewDecoder(bytes.NewReader(data))
	config := new(Config)
	err = dec.Decode(config)
	check(err)
	return config
}
