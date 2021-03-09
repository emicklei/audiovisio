package main

import "time"

type Slide struct {
	ID    string `yaml:"id"`
	Title string `yaml:"title"`
	Image string `yaml:"image"`
	Sound string `yaml:"sound"`

	PauseBeforeAudio time.Duration `yaml:"pause-before-audio"`
	PauseAfterAudio  time.Duration `yaml:"pause-after-audio"`

	NextID string `yaml:"next"`
}
