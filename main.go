package main

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	oConfig = flag.String("c", "test.yaml", "configuration YAML file")
	oOutput = flag.String("o", "page.html", "output HTML file")
)

func main() {
	flag.Parse()
	config := loadConfig()

	page := Page{
		Config:       config,
		Slides:       config.Slides,
		FirstSlideID: config.Leader.NextID,
	}
	log.Printf("detected %d slides\n", len(page.Slides))

	outputDir := filepath.Dir(*oConfig)
	if len(config.JSInclude) == 0 {
		config.JSInclude = "audiovisio.js"
		js := filepath.Join(outputDir, config.JSInclude)
		if !fileExists(js) {
			ioutil.WriteFile(js, []byte(javascriptSource()), os.ModePerm)
		} else {
			log.Println("skip overwriting", js)
		}
	} else {
		log.Println("skip writing", config.JSInclude)
	}
	if len(config.CSSInclude) == 0 {
		config.CSSInclude = "audiovisio.css"
		css := filepath.Join(outputDir, config.CSSInclude)
		if !fileExists(css) {
			ioutil.WriteFile(css, []byte(stylesheetSource()), os.ModePerm)
		} else {
			log.Println("skip overwriting", css)
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

	log.Println("writing", *oOutput)
	ioutil.WriteFile(*oOutput, buf.Bytes(), os.ModePerm)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func check(err error) {
	if err == nil {
		return
	}
	log.Fatalln(err)
}
