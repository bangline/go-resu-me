package resume

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonalInfoParsing(t *testing.T) {
	resume, err := LoadResume("testdata/personal_info.yml")
	if err != nil {
		t.Error("failed loading yaml", err)
	}
	assert.Equal(t, "David Kennedy", resume.Name)
	assert.Equal(t, "Glasgow, UK", resume.Details[0])
	assert.Equal(t, "david@bangline.co.uk", resume.Details[1])
	assert.Equal(t, "+44(0)123456789", resume.Details[2])
	assert.Equal(t, "[my website](http://mywebsite.com)", resume.Details[3])
}

func TestBasicSectionParsing(t *testing.T) {
	resume, err := LoadResume("testdata/basic_section.yml")
	if err != nil {
		t.Error("failed loading yaml", err)
	}
	assert.Equal(t, 2, len(resume.Sections))
	assert.Equal(t, "My first section", resume.Sections[0].Title)
	assert.Equal(t, "My second section", resume.Sections[1].Title)

	if entries, ok := resume.Sections[0].Entries.([]string); ok {
		assert.Equal(t, 2, len(entries))
		assert.Equal(t, "entry 1", entries[0])
		assert.Equal(t, "entry 2", entries[1])
	} else {
		t.Error("Expected first section entries to be []string")
	}

	if entries, ok := resume.Sections[1].Entries.([]string); ok {
		assert.Equal(t, 3, len(entries))
		assert.Equal(t, "entry 3", entries[0])
		assert.Equal(t, "entry 4", entries[1])
		assert.Equal(t, "entry 5", entries[2])
	} else {
		t.Error("Expected first section entries to be []string")
	}
}

func TestWorkSectionParsing(t *testing.T) {
	resume, err := LoadResume("testdata/work_section.yml")
	if err != nil {
		t.Error("failed loading yaml", err)
	}
	assert.Equal(t, 1, len(resume.Sections))
	if entries, ok := resume.Sections[0].Entries.([]WorkEntry); ok {
		assert.Equal(t, 2, len(entries))
		assert.Equal(t, "Robot 1", entries[0].Position)
		assert.Equal(t, "Acme", entries[0].Company)
		assert.Equal(t, "Summer - Winter", entries[0].Dates)
		assert.Equal(t, 2, len(entries[0].Bullets))
		assert.Equal(t, "bullet 1", entries[0].Bullets[0])
		assert.Equal(t, "bullet 2", entries[0].Bullets[1])

		assert.Equal(t, "Blahmo", entries[1].Company)
		assert.Equal(t, "May - December", entries[1].Dates)
		assert.Equal(t, 2, len(entries[1].Roles))
		assert.Equal(t, "Robot 2", entries[1].Roles[0].Position)
		assert.Equal(t, "November - December", entries[1].Roles[0].Dates)
		assert.Equal(t, "A-Team", entries[1].Roles[0].Team)
		assert.Equal(t, 3, len(entries[1].Roles[0].Bullets))
	}
}
