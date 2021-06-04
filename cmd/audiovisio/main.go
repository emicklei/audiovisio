package main

import (
	"flag"
	"path/filepath"

	"github.com/emicklei/audiovisio"
)

var (
	oConfig = flag.String("c", "test.yaml", "configuration YAML file")
	oOutput = flag.String("o", "page.html", "output HTML file")
)

func main() {
	flag.Parse()
	config := audiovisio.LoadConfig(*oConfig)
	outputDir := filepath.Dir(*oConfig)
	audiovisio.BuildPresentation(config, outputDir)
}
