package db

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kalpit-sharma-dev/url-shortner/src/constants"
	"github.com/kalpit-sharma-dev/url-shortner/src/model"
	"github.com/kalpit-sharma-dev/url-shortner/src/repository"
)

type FileImpl struct {
	fileProvider *os.File
}

// NewRepository returns instance of File repository
func NewFileRepository(file *os.File) repository.FileRepository {

	return &FileImpl{fileProvider: file}
}

// CreateUrl implements repository.FileRepository
func (file *FileImpl) CreateUrl(ctx context.Context, req model.Request) (urlResp model.Response, err error) {
	fmt.Println("3333333333333333333333333333333333333333333", req)
	writer := csv.NewWriter(file.fileProvider)
	line := []string{strconv.Itoa(req.ID) + constants.Delimeter + string(req.Url) + constants.Delimeter + string(req.Domain)}
	fmt.Println("444444444444444444444444444444444444444444444", line)
	err = writer.Write(line)
	if err != nil {
		log.Fatal("Unable to write to csv", err)

	}
	writer.Flush()
	return urlResp, err
	//panic("unimplemented")
}

// GetUrl implements repository.FileRepository
func (file *FileImpl) GetUrl(ctx context.Context, url string) (urlResp model.Response, err error) {
	return urlResp, err
}
