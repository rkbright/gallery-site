package views

import (
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	LayoutDir  string = "views/layouts/"
	TemplatExt string = ".gohtml"
)

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

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {

}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplatExt)
	if err != nil {
		panic(err)
	}
	return files
}
