package resume

import "gopkg.in/yaml.v3"

type Section struct {
	Title   string      `yaml:"title"`
	Entries interface{} `yaml:"entries"`
}

type BasicSection struct {
	Title   string   `yaml:"title"`
	Entries []string `yaml:"entries"`
}

type WorkSection struct {
	Title   string      `yaml:"title"`
	Entries []WorkEntry `yaml:"entries"`
}

type WorkEntry struct {
	Type     string   `yaml:"type"`
	Position string   `yaml:"position,omitempty"`
	Company  string   `yaml:"company"`
	Location string   `yaml:"location"`
	Dates    string   `yaml:"dates"`
	Bullets  []string `yaml:"bullets,omitempty"`
	Roles    []Role   `yaml:"roles,omitempty"`
}

type Role struct {
	Position string   `yaml:"position"`
	Dates    string   `yaml:"dates"`
	Team     string   `yaml:"team"`
	Bullets  []string `yaml:"bullets"`
}

func (s *Section) UnmarshalYAML(value *yaml.Node) error {
	// Temporary struct to parse the basic structure
	var temp struct {
		Title   string        `yaml:"title"`
		Entries []interface{} `yaml:"entries"`
	}

	if err := value.Decode(&temp); err != nil {
		return err
	}

	s.Title = temp.Title

	// If entries is empty, return early
	if len(temp.Entries) == 0 {
		s.Entries = []string{}
		return nil
	}

	// Check the first entry to determine the section type
	firstEntry := temp.Entries[0]

	// Try to type assert as map[string]interface{} to check for work entries
	if mapEntry, ok := firstEntry.(map[string]interface{}); ok {
		// Check if it's a work entry
		if entryType, exists := mapEntry["type"]; exists && entryType == "work" {
			// It's a work section
			var workSection WorkSection
			if err := value.Decode(&workSection); err != nil {
				return err
			}
			s.Entries = workSection.Entries
			return nil
		}
	}

	// If we get here, it's a basic section
	var basicSection BasicSection
	if err := value.Decode(&basicSection); err != nil {
		return err
	}
	s.Entries = basicSection.Entries
	return nil
}
