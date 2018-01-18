package templates

import (
	"text/template"

	"github.com/chtison/libaws/pkg/fileutils"
	"github.com/chtison/libgo/tmpl"
)

type (
	// Template ...
	Template struct {
		*template.Template
		Libaws *Libaws
	}
	// Libaws ...
	Libaws struct {
		S3Bucket       string
		S3Prefix       string
		UploadFileToS3 func(bucket, localFilePath string) (string, error)
		Sha256File     func(filepath string) (string, error)
	}
)

// New ...
func New() *Template {
	laws := NewLibaws()
	return &Template{
		Template: template.Must(template.New("libaws").Funcs(tmpl.Funcs()).Funcs(template.FuncMap{
			"libaws": laws.This,
		}).Parse(Templates)),
		Libaws: laws,
	}
}

// NewLibaws ...
func NewLibaws() *Libaws {
	return &Libaws{
		UploadFileToS3: fileutils.UploadFileToS3,
		Sha256File:     fileutils.Sha256File,
	}
}

// This ...
func (libaws *Libaws) This() *Libaws {
	return libaws
}
