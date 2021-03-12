package main

type Slide struct {
	ID    string `yaml:"id"`
	Title string `yaml:"title"`
	Image string `yaml:"image"`
	Sound string `yaml:"sound"`

	PauseBeforeAudio int `yaml:"pause-before-audio"`
	PauseAfterAudio  int `yaml:"pause-after-audio"`

	NextID string `yaml:"next"`
}

func (s Slide) withPauses(before, after int) Slide {
	if s.PauseBeforeAudio == 0 {
		s.PauseBeforeAudio = before
	}
	if s.PauseAfterAudio == 0 {
		s.PauseAfterAudio = after
	}
	return s
}
