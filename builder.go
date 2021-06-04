package audiovisio

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func BuildPresentation(config *Config, outputDir string) {
	page := Page{
		Config:       config,
		Slides:       config.Slides,
		FirstSlideID: config.Leader.NextID,
	}
	log.Printf("detected %d slides\n", len(page.Slides))

	if len(config.JSInclude) == 0 {
		config.JSInclude = "audiovisio.js"
		js := filepath.Join(outputDir, config.JSInclude)
		if !fileExists(js) {
			log.Println("writing", js)
			ioutil.WriteFile(js, []byte(javascriptSource()), os.ModePerm)
		} else {
			log.Println("skip writing", js)
		}
	} else {
		log.Println("skip writing", config.JSInclude)
	}
	if len(config.CSSInclude) == 0 {
		config.CSSInclude = "audiovisio.css"
		css := filepath.Join(outputDir, config.CSSInclude)
		if !fileExists(css) {
			log.Println("writing", css)
			ioutil.WriteFile(css, []byte(stylesheetSource()), os.ModePerm)
		} else {
			log.Println("skip writing", css)
		}
	} else {
		log.Println("skip writing", config.CSSInclude)
	}

	log.Println("processing HTML template")
	//html, _ := ioutil.ReadFile("html.txt")
	buf := new(bytes.Buffer)
	thtml, _ := template.New("Zlidez").Parse(htmlTemplateSource())
	if err := thtml.ExecuteTemplate(buf, "Zlidez", page); err != nil {
		log.Fatal(err)
	}

	log.Println("writing to", outputDir)
	ioutil.WriteFile(outputDir, buf.Bytes(), os.ModePerm)
}
