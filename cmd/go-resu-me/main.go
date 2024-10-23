package main

import (
	"flag"
	"log"

	"go-resu-me.bangline.dev/internal/resume"
)

func main() {
	sourcePath := flag.String("source", "yaml/master.yml", "Path to yaml file to build resume from")
	outputPath := flag.String("out", "resume.html", "Path to output html file for the resume")
	themeName := flag.String("theme", "default", "Theme to use for the resume")
	flag.Parse()

	resume, err := resume.LoadResume(*sourcePath)
	if err != nil {
		log.Fatal(err)
	}

	err = resume.Generate(*outputPath, *themeName)
	if err != nil {
		log.Fatal(err)
	}
}
