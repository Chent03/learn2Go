package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// View struct for a new View
type View struct {
	Template *template.Template
	Layout   string
}

var (
	// LayoutDir specifies the layout directory
	LayoutDir string = "views/layouts/"
	// TemplateExt specifies the extendsion to look for
	TemplateExt string = ".gohtml"
)

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// NewView appends footers to new view
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// Render handles the execution of the template
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
