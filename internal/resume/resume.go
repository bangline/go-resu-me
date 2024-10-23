package resume

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"go-resu-me.bangline.dev/internal/template_funcs"
	"gopkg.in/yaml.v3"
)

type Resume struct {
	Colors            Colors
	Name              string    `yaml:"name"`
	Details           []string  `yaml:"details"`
	PersonalStatement string    `yaml:"personal_statement"`
	Sections          []Section `yaml:"sections"`
}
type Colors struct {
	Bg      string
	Heading string
	Text    string
	Accent  string
}

func LoadResume(yamlFilePath string) (*Resume, error) {
	var resume Resume
	yml, err := os.ReadFile(yamlFilePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yml, &resume)
	if err != nil {
		return nil, err
	}

	return &resume, nil
}

func (r *Resume) Generate(path string, theme string, colors Colors) error {
	templateFiles := []string{
		"templates/" + theme + "/resume.html.tmpl",
		"templates/" + theme + "/header.html.tmpl",
		"templates/" + theme + "/basic.html.tmpl",
		"templates/" + theme + "/work.html.tmpl",
		"templates/" + theme + "/multiple_roles.html.tmpl",
		"templates/" + theme + "/single_role.html.tmpl",
	}

	r.Colors = colors

	tmpl, err := template.New("resume.html.tmpl").Funcs(templateHelperFunctions()).ParseFiles(templateFiles...)
	if err != nil {
		return fmt.Errorf("error parsing templates: %w", err)
	}

	html, err := os.Create(path)
	if err != nil {
		log.Fatal("could not write to file %w", err)
	}
	defer html.Close()

	return tmpl.Execute(html, r)
}

func templateHelperFunctions() template.FuncMap {
	return template.FuncMap{
		"inPairs":         template_funcs.InPairs,
		"md":              template_funcs.Markdown,
		"isBasicSection":  isBasicSection,
		"isWorkSection":   isWorkSection,
		"getWorkEntries":  getWorkEntries,
		"getBasicEntries": getBasicEntries,
		"hasRoles":        hasRoles,
	}
}

func isBasicSection(section Section) bool {
	_, ok := section.Entries.([]string)
	return ok
}

func isWorkSection(section Section) bool {
	_, ok := section.Entries.([]WorkEntry)
	return ok
}

func getBasicEntries(section Section) []string {
	if entries, ok := section.Entries.([]string); ok {
		return entries
	}
	return nil
}

func getWorkEntries(section Section) []WorkEntry {
	if entries, ok := section.Entries.([]WorkEntry); ok {
		return entries
	}
	return nil
}

func hasRoles(entry WorkEntry) bool {
	return len(entry.Roles) > 0
}
