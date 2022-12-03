package service

import (
	"context"
	"fmt"

	"github.com/kalpit-sharma-dev/url-shortner/src/model"
	"github.com/kalpit-sharma-dev/url-shortner/src/repository"
)

type UrlServiceImpl struct {
	repository repository.FileRepository
}

// CreateUrl implements UrlService
func (s *UrlServiceImpl) CreateUrl(ctx context.Context, urlReq model.Request) (urlResp model.Response, err error) {
	fmt.Println("222222222222222222222222222222222222222222222222222222222222222222222", urlReq)
	s.repository.CreateUrl(ctx, urlReq)
	return urlResp, nil
	//panic("unimplemented")
}

// GetUrl implements UrlService
func (s *UrlServiceImpl) GetUrl(ctx context.Context, url string) (urlResp model.Response, err error) {
	return urlResp, nil
}

func NewUrlService(repository repository.FileRepository) UrlService {
	return &UrlServiceImpl{repository: repository}
}
