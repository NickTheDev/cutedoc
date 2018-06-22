package cutedoc

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// config represents the default name of the cutedoc config file.
const config = ".cutedoc.yml"

// Doc represents the current model of the documentation being created.
type Doc struct {
	Build    Build
	Meta     Meta
	Theme    Theme
	Articles yaml.MapSlice
}

// Build represents options used during the build process.
type Build struct {
	Dir    string
	Minify bool
}

// Meta holds metadata of the documentation.
type Meta struct {
	Title       string
	Author      string
	Description string
	Branding    Branding
}

// Represents theme of the documentation.
type Theme struct {
	Template string
	Colors   Colors
}

// Represents color scheme of the documentation.
type Colors struct {
	Primary    string
	Secondary  string
	Text       string
	Nav        string
	Background string
}

// Branding holds the paths to the branding of the documentation.
type Branding struct {
	Logo string
	Icon string
}

// New reads the doc config in the current directory and creates a doc model based on it.
func New() (*Doc, error) {
	bytes, err := ioutil.ReadFile(config)

	if err != nil {
		return nil, err
	}

	doc := &Doc{Build{"docs", true},
		Meta{"Docs", "", "Documentation page.", Branding{}},
		Theme{"minimal", Colors{"3e4669", "c3c6de", "3e4669", "fff", "f6f7fb"}},
		yaml.MapSlice{}}

	return doc, yaml.Unmarshal(bytes, doc)
}
