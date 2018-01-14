package libaws

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/chtison/libaws/libaws/format"
	"github.com/chtison/libaws/template/templates"
	"github.com/chtison/libgo/cli"
	"github.com/chtison/libgo/fmt"
	yaml "gopkg.in/yaml.v2"
)

// DefaultOutputDir ...
const DefaultOutputDir = "_out"

// FlagConfig ...
var FlagConfig = cli.NewFlagString('c', "config", "libaws.yaml")

func init() {
	FlagConfig.Usage().Synopsys = "Required. Set the libaws configuration file path. Default: libaws.yaml"
}

// WriteFile ...
func WriteFile(filename, content string) error {
	w, err := os.OpenFile(filename, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer w.Close()
	if _, err = io.WriteString(w, content); err != nil {
		return err
	}
	fmt.Printfln(`%s successfully created`, filename)
	return nil
}

// ReadYamlFile ...
func ReadYamlFile(filename string, data interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(b, data); err != nil {
		return err
	}
	return nil
}

// LibAws ...
type LibAws struct {
	Root     *format.Root
	Template *template.Template
	Data     interface{}
}

// NewLibAws ...
func NewLibAws() *LibAws {
	return &LibAws{}
}

// NewLibAwsFromFlagConfig ...
func NewLibAwsFromFlagConfig() (*LibAws, error) {
	libaws := NewLibAws()
	if err := ReadYamlFile(FlagConfig.Value, &libaws.Root); err != nil {
		return nil, err
	}
	return libaws, nil
}

// ReadDataFiles ...
func (libaws *LibAws) ReadDataFiles() error {
	for _, filename := range libaws.Root.Datas {
		if err := ReadYamlFile(filename, &libaws.Data); err != nil {
			return err
		}
	}
	return nil
}

// ReadTemplateFiles ...
func (libaws *LibAws) ReadTemplateFiles() error {
	for _, t := range libaws.Root.Templates {
		name := t.Name
		if name == "" {
			name = t.Path
		}
		if err := libaws.AddTemplate(name, t.Path); err != nil {
			return err
		}
	}
	return nil
}

// AddTemplate ...
func (libaws *LibAws) AddTemplate(name, filename string) error {
	if libaws.Template == nil {
		libaws.Template = template.Must(templates.Template.Clone())
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if _, err = libaws.Template.New(name).Parse(string(b)); err != nil {
		return err
	}
	if libaws.Root.TemplateToExecute == "" {
		libaws.Root.TemplateToExecute = name
	}
	return nil
}

// ZipLambda ...
func (libaws *LibAws) ZipLambda(path string) error {
	if len(libaws.Root.Lambdas) == 0 {
		return fmt.Errorf(`no lambda function to zip`)
	}
	for _, lambda := range libaws.Root.Lambdas {
		lambdaPath := filepath.Clean(lambda.Path)
		if path == "" || lambdaPath == path {
			if err := libaws.zipLambda(lambdaPath, lambda); err != nil {
				return err
			}
			if path == "" {
				continue
			}
			return nil
		}
	}
	if path != "" {
		return fmt.Errorf(`lambda function not found for path: "%s"`, path)
	}
	return nil
}

func (libaws *LibAws) zipLambda(lambdaPath string, lambda format.Lambda) error {
	outPath, err := libaws.getOutputLambdaZipFilePath(lambdaPath, lambda)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	z := zip.NewWriter(buf)
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil || path == lambdaPath || info.IsDir() {
			return err
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		path, err = filepath.Rel(lambdaPath, path)
		if err != nil {
			return err
		}
		zipHeader := &zip.FileHeader{Name: path}
		zipHeader.SetMode(0600)
		w, err := z.CreateHeader(zipHeader)
		if err != nil {
			return err
		}
		if _, err := io.Copy(w, f); err != nil {
			return err
		}
		return nil
	}
	f, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_SYNC, 0700)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = filepath.Walk(lambdaPath, walkFunc); err != nil {
		return err
	}
	if err = z.Close(); err != nil {
		return err
	}
	if _, err = io.Copy(f, buf); err != nil {
		return err
	}
	return nil
}

func (libaws *LibAws) getOutputLambdaZipFilePath(path string, lambda format.Lambda) (string, error) {
	outDir := lambda.OutDir
	if outDir == "" {
		outDir = libaws.Root.OutDir
		if outDir == "" {
			outDir = filepath.Join(DefaultOutputDir, "lambda")
		}
	}
	if err := os.MkdirAll(outDir, 0700); err != nil {
		return "", err
	}
	outFile := lambda.OutFile
	if outFile == "" {
		outFile = strings.TrimSuffix(lambda.Path, filepath.Ext(lambda.Path)) + ".zip"
	}
	return filepath.Clean(filepath.Join(outDir, outFile)), nil
}
