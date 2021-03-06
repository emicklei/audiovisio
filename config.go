package audiovisio

import (
	"bytes"
	"html/template"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Title            string        `yaml:"title"`
	CSSInclude       string        `yaml:"css"`
	JSInclude        string        `yaml:"js"`
	HeaderContent    template.HTML `yaml:"header"`
	FooterContent    template.HTML `yaml:"footer"`
	ButtonStartTitle string        `yaml:"button-start-title"`

	PauseBeforeAudio int `yaml:"pause-before-audio"`
	PauseAfterAudio  int `yaml:"pause-after-audio"`

	Leader  Slide   `yaml:"leader"`
	Slides  []Slide `yaml:"slides"`
	Trailer Slide   `yaml:"trailer"`
}

func (c *Config) postLoad() {
	// propagate the pauses
	changed := []Slide{}
	for _, each := range c.Slides {
		changed = append(changed, each.withPauses(c.PauseBeforeAudio, c.PauseAfterAudio))
	}
	c.Slides = changed
	if len(c.ButtonStartTitle) == 0 {
		c.ButtonStartTitle = "Start"
	}
}

func LoadConfig(cfg string) *Config {
	data, err := ioutil.ReadFile(cfg)
	check(err)
	dec := yaml.NewDecoder(bytes.NewReader(data))
	config := new(Config)
	err = dec.Decode(config)
	check(err)
	config.postLoad()
	return config
}
