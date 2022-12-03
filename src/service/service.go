package service

import (
	"context"

	"github.com/kalpit-sharma-dev/url-shortner/src/model"
	"github.com/kalpit-sharma-dev/url-shortner/src/repository"
)

type UrlServiceImpl struct {
	repository repository.FileRepository
}

// CreateUrl implements UrlService
func (s *UrlServiceImpl) CreateUrl(ctx context.Context, urlReq model.Request) (urlResp model.Response, err error) {
	urlResp, err = s.repository.GetUrl(ctx, urlReq.Url)
	if err != nil {
		return
	}
	if len(urlResp.Url) != 0 {
		return
	}
	tinyUrl := randStringBytesRmndr(7)

	urlResp, err = s.repository.CreateUrl(ctx, urlReq.Url, tinyUrl)
	return urlResp, err
}

// GetUrl implements UrlService
func (s *UrlServiceImpl) GetUrl(ctx context.Context, url string) (urlResp model.Response, err error) {
	urlResp, err = s.repository.GetUrl(ctx, url)
	return
}

func NewUrlService(repository repository.FileRepository) UrlService {
	return &UrlServiceImpl{repository: repository}
}
