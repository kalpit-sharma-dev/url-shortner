package repository

import (
	"context"

	"github.com/kalpit-sharma-dev/url-shortner/src/model"
)

type FileRepository interface {
	CreateUrl(ctx context.Context, url, tinyUrl string) (urlResp model.Response, err error)
	GetUrl(ctx context.Context, url string) (urlResp model.Response, err error)
}
