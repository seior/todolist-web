package view

import (
	"embed"
	"text/template"
)

//go:embed *.gohtml
var file embed.FS

var ViewHTML = template.Must(template.ParseFS(file, "*.gohtml"))