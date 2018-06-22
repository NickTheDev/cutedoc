package cutedoc

import (
	"bytes"
	"errors"
	"github.com/nickthedev/cutedoc/template"
	"github.com/shurcooL/github_flavored_markdown"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/js"
	"io/ioutil"
	"log"
	"os"
	"strings"
	tmpl "text/template"
)

var (
	// tasks represents all the steps of the generation process.
	tasks = []task{
		initBuild,
		copyBranding,
		copyScript,
		copyStatic,
		copyTemplate,
	}

	// templates represent the supported themes.
	templates = map[string]string{
		"minimal": "template/minimal",
	}

	// packer compresses assets for production.
	minifier = minify.New()
)

// tasks represents a step of the generation process.
type task func(doc *Doc) error

// copyImage copies the image to the output minimal directory if it is not null.
func copyImage(doc *Doc, image, output string) error {
	if image != "" {
		icon, err := ioutil.ReadFile(image)

		if err != nil {
			return err
		}

		if err := os.MkdirAll(doc.Build.Dir+"/static/img", os.ModePerm); err != nil {
			return err
		}

		ioutil.WriteFile(doc.Build.Dir+"/static/img/"+output, icon, os.ModePerm)
	}

	return nil
}

// initBuild checks the template and creates the output directory.
func initBuild(doc *Doc) error {
	if _, ok := templates[doc.Theme.Template]; !ok {
		return errors.New("unsupported template theme")
	}

	if err := os.RemoveAll(doc.Build.Dir); err != nil {
		return err
	}

	return os.MkdirAll(doc.Build.Dir, os.ModePerm)
}

// copyBranding copies the logo and icon if they are not null to the minimal output directory.
func copyBranding(doc *Doc) error {
	if err := copyImage(doc, doc.Meta.Branding.Icon, "icon.png"); err != nil {
		return err
	}

	if err := copyImage(doc, doc.Meta.Branding.Logo, "logo.png"); err != nil {
		return err
	}

	return nil
}

// copyScript copies the template script output minimal directory.
func copyScript(doc *Doc) error {
	if err := os.MkdirAll(doc.Build.Dir+"/static/js", os.ModePerm); err != nil {
		return err
	}

	asset, err := template.Asset(templates[doc.Theme.Template] + "/main.js.tmpl")

	if err != nil {
		return err
	}

	if doc.Build.Minify {
		packed, err := minifier.Bytes("text/javascript", asset)

		if err != nil {
			return err
		} else {
			asset = packed
		}
	}

	return ioutil.WriteFile(doc.Build.Dir+"/static/js/main.js", asset, os.ModePerm)
}

// copyStatic transfers the template stylesheet to the output minimal directory.
func copyStatic(doc *Doc) error {
	if err := os.MkdirAll(doc.Build.Dir+"/static/css", os.ModePerm); err != nil {
		return err
	}

	asset, err := template.Asset(templates[doc.Theme.Template] + "/main.css.tmpl")

	if err != nil {
		return err
	}

	format, err := tmpl.New("style").Parse(string(asset))

	if err != nil {
		return err
	}

	var buffer bytes.Buffer

	if err := format.Execute(&buffer, doc.Theme.Colors); err != nil {
		return err
	}

	var data = buffer.Bytes()

	if doc.Build.Minify {
		data, err = minifier.Bytes("text/css", data)

		if err != nil {
			return err
		}
	}

	return ioutil.WriteFile(doc.Build.Dir+"/static/css/main.css", data, os.ModePerm)
}

// copyTemplate injects the doc model into the template page and copies it to the output directory.
func copyTemplate(doc *Doc) error {
	asset, err := template.Asset(templates[doc.Theme.Template] + "/index.html.tmpl")

	if err != nil {
		return err
	}

	format, err := tmpl.New("doc").Funcs(tmpl.FuncMap{
		"readContent": func(path string) string {
			content, err := ioutil.ReadFile(path)

			if err != nil {
				log.Fatalf("An error occured generating documentation: \n\t%v.", err)
			}

			return string(github_flavored_markdown.Markdown(content))
		},
		"idFrom": func(name string) string {
			return strings.Replace(strings.ToLower(name), " ", "-", -1)
		},
	}).Parse(string(asset))

	if err != nil {
		return err
	}

	dst, err := os.Create(doc.Build.Dir + "/index.html")

	if err != nil {
		return err
	}

	defer dst.Close()

	return format.Execute(dst, doc)
}

// Run executes the documentation generation tasks.
func Run(doc *Doc) error {
	if doc.Build.Minify {
		minifier.AddFunc("text/css", css.Minify)
		minifier.AddFunc("text/javascript", js.Minify)
	}

	for _, task := range tasks {
		if err := task(doc); err != nil {
			return err
		}

	}

	return nil
}
