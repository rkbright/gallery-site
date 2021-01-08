package views

import (
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

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplatExt)
	if err != nil {
		panic(err)
	}
	return files
}
