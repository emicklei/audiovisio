package main

type Slide struct {
	ID    string `yaml:"id"`
	Title string `yaml:"title"`
	Image string `yaml:"image"`
	Sound string `yaml:"sound"`

	NextID string `yaml:"next"`
}
