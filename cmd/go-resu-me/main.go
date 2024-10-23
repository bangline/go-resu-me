package main

import (
	"flag"
	"log"

	"go-resu-me.bangline.dev/internal/resume"
)

func main() {
	sourcePath := flag.String("source", "resume.yml", "Path to yaml file to build resume from")
	outputPath := flag.String("out", "resume.html", "Path to output html file for the resume")
	themeName := flag.String("theme", "default", "Theme to use for the resume")
	bgColor := flag.String("bg", "#f5f3ff", "Color of the web background")
	headingColor := flag.String("h", "#0a0a0a", "Color of the heading text")
	textColor := flag.String("t", "#0c0a09", "Color of the non heading text")
	accentColor := flag.String("a", "#2e1065", "Color of the accent text")
	flag.Parse()

	colors := resume.Colors{
		Bg:      *bgColor,
		Heading: *headingColor,
		Text:    *textColor,
		Accent:  *accentColor,
	}

	resume, err := resume.LoadResume(*sourcePath)
	if err != nil {
		log.Fatal(err)
	}

	err = resume.Generate(*outputPath, *themeName, colors)
	if err != nil {
		log.Fatal(err)
	}
}
