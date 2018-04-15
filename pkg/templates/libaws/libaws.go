package libaws

import (
	"github.com/chtison/libgo/fmt"
	"github.com/chtison/libgo/tmpl/pkg/utils"
	"text/template"
)

type Libaws struct{
	Template *template.Template
}

func (*Libaws) LogicalId(logicalIdPrefix string, m utils.Map) string {
	return fmt.Sprintf("%s%s", logicalIdPrefix, m.GetDefault("", "LogicalIdSuffix"))
}

func Funcs(t *template.Template) template.FuncMap {
	return template.FuncMap{
		"libaws": func() *Libaws { return &Libaws{t} },
	}
}

type builder struct{
	*fmt.Builder
}

func newBuilder() *builder {
	return &builder{
		Builder: fmt.NewBuilder(),
	}
}

func (b *builder) writeString(s string, err error) error {
	if err != nil {
		return err
	}
	_, err = b.WriteString(s)
	return err
}
