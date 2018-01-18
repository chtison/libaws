package fileutils

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"github.com/chtison/libgo/fmt"
	yaml "gopkg.in/yaml.v2"
)

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

// UploadFileToS3 ...
func UploadFileToS3(bucket, localFilePath string) (string, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return "", err
	}
	uploader := s3manager.NewUploader(cfg)
	f, err := os.Open(localFilePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	keyPath, err := sha256(f)
	if err != nil {
		return "", err
	}
	if _, err = f.Seek(0, io.SeekStart); err != nil {
		return "", err
	}
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(keyPath),
		Body:   f,
	})
	if err != nil {
		return "", err
	}
	return result.Location, nil
}

// Sha256File ...
func Sha256File(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return sha256(f)
}

// sha256 ...
func sha256(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	sum := sha512.Sum512_256(b)
	return hex.EncodeToString(sum[:]), nil
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
