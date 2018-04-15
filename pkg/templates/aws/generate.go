// +build ignore

package main

import (
	"container/list"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/chtison/libgo/fmt"
)

type CloudFormation struct {
	ResourceSpecificationVersion string
	ResourceTypes                map[string]*Resource
	PropertyTypes                map[string]*Resource
}

type Resource struct {
	Properties map[string]*Property
}

type Property struct {
	Name              string `json:"-"`
	ItemType          string
	PrimitiveItemType string
	PrimitiveType     string
	Type              string
	Required          bool
}

type Properties struct {
	Optional []*Property
	Required []*Property
}

func NewProperties(p map[string]*Property) *Properties {
	pp := &Properties{}
	for k, v := range p {
		v.Name = k
		if v.Required {
			pp.Required = append(pp.Required, v)
		} else {
			pp.Optional = append(pp.Optional, v)
		}
	}
	if pp.Optional != nil {
		sort.Slice(pp.Optional, func(i, j int) bool { return pp.Optional[i].Name < pp.Optional[j].Name })
	}
	if pp.Required != nil {
		sort.Slice(pp.Required, func(i, j int) bool { return pp.Required[i].Name < pp.Required[j].Name })
	}
	return pp
}

func main() {
	b, err := ioutil.ReadFile("_tmp/CloudFormationResourceSpecification.json")
	if err != nil {
		fatal(err)
	}
	cf := &CloudFormation{}
	if err := json.Unmarshal(b, &cf); err != nil {
		fatal(err)
	}
	// TO BE REMOVED WHEN OFFICIALY FIXED
	for {
		p := cf.PropertyTypes["Tag"].Properties
		p["Key"].Required = true
		p["Value"].Required = true
		break
	}
	if err := cf.Generate(); err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Fprintfln(os.Stderr, "Error: %s", err)
	os.Exit(1)
}

func (cf *CloudFormation) Generate() error {
	fmt.Println("CloudFormation version:", cf.ResourceSpecificationVersion)
	for resourceTypeName := range cf.ResourceTypes {
		if strings.HasPrefix(resourceTypeName, "AWS::EMR") {
			continue
		}
		tb := &TemplateBuilder{
			CloudFormation:   cf,
			Builder:          &strings.Builder{},
			ResourceTypeName: resourceTypeName,
			BaseIndent:       "    ",
		}
		if err := tb.GenerateTemplate(); err != nil {
			return err
		}
		if err := tb.CreateFile(); err != nil {
			return err
		}
	}
	return nil
}

type TemplateBuilder struct {
	*CloudFormation
	Builder          *strings.Builder
	ResourceTypeName string
	BaseIndent       string
	Indent           string
	Variable         *list.List
}

func (tb *TemplateBuilder) Printf(format string, a ...interface{}) {
	fmt.Fprintf(tb.Builder, format, a...)
}

func (tb *TemplateBuilder) Printif(format string, a ...interface{}) {
	fmt.Fprint(tb.Builder, tb.Indent)
	fmt.Fprintf(tb.Builder, format, a...)
}

func (tb *TemplateBuilder) Printfln(format string, a ...interface{}) {
	fmt.Fprintfln(tb.Builder, format, a...)
}

func (tb *TemplateBuilder) Printifln(format string, a ...interface{}) {
	fmt.Fprintf(tb.Builder, tb.Indent)
	fmt.Fprintfln(tb.Builder, format, a...)
}

func (tb *TemplateBuilder) GenerateTemplate() error {
	tb.Indent = tb.BaseIndent
	tb.Variable = list.New()
	tb.Variable.PushFront("")
	tb.Printfln(`{{- define "%s" }}`, tb.ResourceTypeName)
	tb.Printfln(`{{- with .LogicalId }}`)
	tb.Printifln(`{{ . }}:`)
	tb.Printfln(`{{- end }}`)
	tb.Indent += tb.BaseIndent
	tb.Printfln(`{{- with .DependsOn }}`)
	tb.Printifln(`DependsOn:`)
	tb.Printfln(`{{- range $_, $v := . }}`)
	tb.Printifln(`%s- {{ $v }}`,tb.BaseIndent)
	tb.Printfln(`{{- end }}`)
	tb.Printfln(`{{- end }}`)
	tb.Printifln("Type: %s", tb.ResourceTypeName)
	p := NewProperties(tb.ResourceTypes[tb.ResourceTypeName].Properties)
	if len(p.Required) > 0 {
		tb.Printifln("Properties:")
		tb.Indent += tb.BaseIndent
		tb.GenerateTemplateProperties(p)
	} else if len(p.Optional) > 0 {
		tb.Printf("{{- if ")
		for i, v := range p.Optional {
			if i > 0 {
				tb.Printf("| or ")
			}
			tb.Printf(".%s ", v.Name)
		}
		tb.Printfln("}}")
		tb.Printifln("Properties:")
		tb.Indent += tb.BaseIndent
		tb.GenerateTemplateProperties(p)
		tb.Printfln("{{- end }}")
	}
	tb.Printfln("{{- end }}")
	return nil
}

func (tb *TemplateBuilder) GenerateTemplateProperties(p *Properties) {
	for _, property := range p.Required {
		tb.GenerateTemplateProperty(property)
	}
	for _, property := range p.Optional {
		tb.GenerateTemplateProperty(property)
	}
}

func (tb *TemplateBuilder) GenerateTemplateProperty(p *Property) {
	if !p.Required {
		tb.Printfln("{{- if %s.%s }}", tb.Variable.Front().Value, p.Name)
	}
	tb.Printif("%s:", p.Name)
	if tb.GeneratePackageableTemplate(p) {
		defer tb.Printfln("{{- end }}")
	}
	if p.PrimitiveType != "" {
		if p.PrimitiveType == "Json" {
			tb.Printfln("")
			tb.Indent += tb.BaseIndent
			tb.Printifln(`{{ yaml.MarshalIndent %s.%s "%s" "%s" | fmt.Sprintf "%%s" }}`, tb.Variable.Front().Value, p.Name, tb.Indent, tb.BaseIndent)
			tb.Indent = strings.TrimSuffix(tb.Indent, tb.BaseIndent)
		} else {
			tb.Printfln(" {{ %s.%s }}", tb.Variable.Front().Value, p.Name)
		}
	} else if p.Type == "List" {
		tb.Printfln("")
		if p.PrimitiveItemType != "" {
			tb.Printfln("{{- range $_, $v := %s.%s }}", tb.Variable.Front().Value, p.Name)
			tb.Printifln("%s- {{ $v }}", tb.BaseIndent)
			tb.Printfln("{{- end }}")
		} else if p.ItemType != "" {
			if p.ItemType == "Tag" {
				if _, ok := tb.PropertyTypes[tb.ResourceTypeName+"."+p.ItemType]; !ok {
					tb.PropertyTypes[tb.ResourceTypeName+"."+p.ItemType] = tb.PropertyTypes["Tag"]
				}
			}
			tb.Printfln("{{- range $_, $v := %s.%s }}", tb.Variable.Front().Value, p.Name)
			tb.Printifln("%s-", tb.BaseIndent)
			tb.Indent += tb.BaseIndent + "  "
			tb.Variable.PushFront("$v")
			tb.GenerateTemplateProperties(NewProperties(tb.PropertyTypes[tb.ResourceTypeName+"."+p.ItemType].Properties))
			tb.Variable.Remove(tb.Variable.Front())
			tb.Indent = strings.TrimSuffix(tb.Indent, tb.BaseIndent+"  ")
			tb.Printfln("{{- end }}")
		} else {
			panic(fmt.Sprintf("%s.%s has no PrimitiveItemType nor ItemType", tb.ResourceTypeName, p.Name))
		}
	} else if p.Type == "Map" {
		tb.Printfln("")
		if p.PrimitiveItemType != "" {
			tb.Printfln("{{- range $k, $v := %s.%s }}", tb.Variable.Front().Value, p.Name)
			tb.Printifln("%s{{ $k }}: {{ $v }}", tb.BaseIndent)
			tb.Printfln("{{- end }}")
		} else if p.ItemType != "" {
			tb.Printfln("{{- range $k, $v := %s.%s }}", tb.Variable.Front().Value, p.Name)
			tb.Printifln("%s{{ $k }}:", tb.BaseIndent)
			tb.Indent += tb.BaseIndent + tb.BaseIndent
			tb.Variable.PushFront("$v")
			tb.GenerateTemplateProperties(NewProperties(tb.PropertyTypes[tb.ResourceTypeName+"."+p.ItemType].Properties))
			tb.Variable.Remove(tb.Variable.Front())
			tb.Indent = strings.TrimSuffix(tb.Indent, tb.BaseIndent+tb.BaseIndent)
			tb.Printfln("{{- end }}")
		} else {
			panic(fmt.Sprintf("%s.%s has no PrimitiveItemType nor ItemType", tb.ResourceTypeName, p.Name))
		}
	} else if p.Type != "" {
		tb.Printfln("")
		tb.Indent += tb.BaseIndent
		tb.Variable.PushFront(tb.Variable.Front().Value.(string) + "." + p.Name)
		tb.GenerateTemplateProperties(NewProperties(tb.PropertyTypes[tb.ResourceTypeName+"."+p.Type].Properties))
		tb.Variable.Remove(tb.Variable.Front())
		tb.Indent = strings.TrimSuffix(tb.Indent, tb.BaseIndent)
	} else {
		panic(fmt.Sprintf("%s.%s has no PrimitiveType nor Type", tb.ResourceTypeName, p.Name))
	}
	if !p.Required {
		tb.Printfln("{{- end }}")
	}
}

func (tb *TemplateBuilder) GeneratePackageableTemplate(p *Property) bool {
	if tb.ResourceTypeName == "AWS::ApiGateway::RestApi" && p.Name == "" {
		tb.Printf(`{{ if istype .BodyS3Location "string" }} {{ .BodyS3Location }}{{ else }}`)
	} else if tb.ResourceTypeName == "AWS::Lambda::Function" && p.Name == "Code" {
		tb.Printf(`{{ if istype .Code "string" }} {{ .Code }}{{ else }}`)
	} else if tb.ResourceTypeName == "AWS::ElasticBeanstalk::ApplicationVersion" && p.Name == "SourceBundle" {
		tb.Printf(`{{ if istype .SourceBundle "string" }} {{ .SourceBundle }}{{ else }}`)
	} else if tb.ResourceTypeName == "AWS::CloudFormation::Stack" && p.Name == "TemplateURL" {
		tb.Printf(`{{ if istype .TemplateURL "string" }} {{ .TemplateURL }}{{ else }}`)
	} else {
		return false
	}
	return true
}

func (tb *TemplateBuilder) CreateFile() error {
	split := strings.Split(tb.ResourceTypeName, "::")
	path := filepath.Join("_tmp", "generated", split[1])
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	path = filepath.Join(path, split[2]+".tmpl.yaml")
	if err := ioutil.WriteFile(path, []byte(tb.Builder.String()), 0644); err != nil {
		return err
	}
	return nil
}
