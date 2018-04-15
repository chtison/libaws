package templates

import (
	"text/template"

	"github.com/chtison/libaws/pkg/templates/aws"
	"github.com/chtison/libaws/pkg/templates/libaws"
	"github.com/chtison/libgo/tmpl"
)

// New ...
func New() *template.Template {
	t := template.New("libaws")
	template.Must(t.Funcs(tmpl.Funcs()).Parse(aws.Templates))
	template.Must(t.Funcs(libaws.Funcs(t)).Parse(libaws.Templates))
	return t
}
