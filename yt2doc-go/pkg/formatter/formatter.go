package formatter

import (
	"os"
	"text/template"

	"github.com/samar-108/yt2doc-go/internal/models"
)

type Formatter struct {
	template *template.Template
}

// NewFormatter parses the markdown template once and returns a new Formatter.
func NewFormatter() (*Formatter, error) {
	tmpl, err := template.New("markdown").Parse(markdownTemplate)
	if err != nil {
		return nil, err
	}
	return &Formatter{template: tmpl}, nil
}

const markdownTemplate = `
# {{.MetaData.Title}}

**Published:** {{.MetaData.PublishedAt.Format "2006-01-02"}}
**Author:** {{.MetaData.Author}}
**Description:** {{.MetaData.Description}}

---

{{range .Chapters}}

## {{.Title}}

{{range .Segments}}
- {{.Text}}
{{end}}

---
{{end}}
`

// ToMarkdown generates a markdown file from the given document.
func (f *Formatter) ToMarkdown(doc *models.Document, outputPath string) error {
	// 1. Create the output file
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 2. Execute the pre-parsed template and write to the file
	return f.template.Execute(file, doc)
}
