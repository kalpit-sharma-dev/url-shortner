package service

import (
	"context"

	"github.com/kalpit-sharma-dev/url-shortner/src/model"
)

type UrlService interface {
	CreateUrl(ctx context.Context, urlReq model.Request) (urlResponse model.Response, err error)
	GetUrl(ctx context.Context, url string) (urlResponse model.Response, err error)
}
