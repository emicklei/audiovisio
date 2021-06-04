package audiovisio

import (
	"log"
	"os"
)

func check(err error) {
	if err == nil {
		return
	}
	log.Fatalln(err)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
