package staticParts

import (
	"bytes"
	"html/template"
	"log"
)

type Breadcrumbs struct {
	Href, Name string
	Active     bool
}

const bcPath string = "project/view/staticPages/parts/breadcrumbs.html"

func GenerateBreadcrumbsHtml(bc Breadcrumbs) template.HTML {
	tmpl, err := template.ParseFiles(bcPath)
	if err != nil {
		log.Printf("Failed to parse template: %v", err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, bc)
	if err != nil {
		log.Printf("Failed to render template: %v", err)
	}

	return template.HTML(buf.String())
}
