package render

import (
	"bytes"
	"fmt"
	"html/template"
)

type Document struct {
	Title        string
	PageNotFound bool
}

func (d Document) Render(t *template.Template) ([]byte, error) {
	return render(t, "document.html", d)
}

func render(t *template.Template, name string, data any) ([]byte, error) {
	var b bytes.Buffer
	if err := t.ExecuteTemplate(&b, name, data); err != nil {
		return nil, fmt.Errorf("executing template %q on %T: %w", name, data, err)
	}
	return b.Bytes(), nil
}
