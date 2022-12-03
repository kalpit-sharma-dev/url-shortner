package db

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/kalpit-sharma-dev/url-shortner/src/constants"
	"github.com/kalpit-sharma-dev/url-shortner/src/model"
	"github.com/kalpit-sharma-dev/url-shortner/src/repository"
)

type FileImpl struct {
	fileProvider *os.File
}

// Read implements io.Reader
func (*FileImpl) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}

// NewRepository returns instance of File repository
func NewFileRepository(file *os.File) repository.FileRepository {

	return &FileImpl{fileProvider: file}
}

// CreateUrl implements repository.FileRepository
func (file *FileImpl) CreateUrl(ctx context.Context, url, tinyUrl string) (urlResp model.Response, err error) {
	writer := csv.NewWriter(file.fileProvider)
	line := []string{url, constants.PrefixUrl + tinyUrl}
	err = writer.Write(line)
	if err != nil {
		log.Fatal("Unable to write to csv", err)
		return

	}
	writer.Flush()
	file.fileProvider.Seek(0, io.SeekStart)
	urlResp.Url = url
	urlResp.TinyUrl = constants.PrefixUrl + tinyUrl
	return urlResp, err
}

// GetUrl implements repository.FileRepository
func (file *FileImpl) GetUrl(ctx context.Context, url string) (urlResp model.Response, err error) {

	_, err = file.fileProvider.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
		return
	}
	reader := csv.NewReader(file.fileProvider)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("unable to read data from csv", err)
		return
	}
	for _, record := range records {

		if record[0] == url {
			urlResp.Url = record[0]
			urlResp.TinyUrl = record[1]
			return
		}
	}
	return
}
