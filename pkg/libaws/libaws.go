package libaws

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/chtison/libaws/pkg/config"
	"github.com/chtison/libaws/pkg/fileutils"
	"github.com/chtison/libaws/pkg/templates"
	"github.com/chtison/libgo/cli"
	yaml "gopkg.in/yaml.v2"
)

// DefaultOutputDir ...
const DefaultOutputDir = "_out"

// FlagConfig ...
var FlagConfig = cli.NewFlagString('c', "config", "libaws.yaml")

func init() {
	FlagConfig.Usage().Synopsys = "Required. Set the libaws configuration file path. Default: libaws.yaml"
}

// Libaws ...
type Libaws struct {
	Root     *config.Root
	Template *templates.Template
	Data     map[string]interface{}
}

// NewLibaws ...
func NewLibaws() *Libaws {
	return &Libaws{}
}

// NewLibawsFromFlagConfig ...
func NewLibawsFromFlagConfig() (*Libaws, error) {
	b, err := ioutil.ReadFile(FlagConfig.Value)
	if err != nil {
		return nil, err
	}
	libaws := NewLibaws()
	if err := yaml.Unmarshal(b, &libaws.Root); err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(b, &libaws.Data); err != nil {
		return nil, err
	}
	return libaws, nil
}

// ReadDataFiles ...
func (libaws *Libaws) ReadDataFiles() error {
	for _, filename := range libaws.Root.Datas {
		if err := fileutils.ReadYamlFile(filename, &libaws.Data); err != nil {
			return err
		}
	}
	return nil
}

// ReadTemplateFiles ...
func (libaws *Libaws) ReadTemplateFiles() error {
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
func (libaws *Libaws) AddTemplate(name, filename string) error {
	if libaws.Template == nil {
		libaws.ReloadTemplate()
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

// ReloadTemplate ...
func (libaws *Libaws) ReloadTemplate() {
	if libaws.Template == nil {
		libaws.Template = templates.New()
	}
	if libaws.Root.S3Bucket != "" {
		libaws.Template.Libaws.S3Bucket = libaws.Root.S3Bucket
	}
	if libaws.Root.S3Prefix != "" {
		libaws.Template.Libaws.S3Prefix = libaws.Root.S3Prefix
	}
}

// ZipLambda ...
func (libaws *Libaws) ZipLambda(path string) error {
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

func (libaws *Libaws) zipLambda(lambdaPath string, lambda *config.Lambda) error {
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

func (libaws *Libaws) getOutputLambdaZipFilePath(path string, lambda *config.Lambda) (string, error) {
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
