package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kalpit-sharma-dev/url-shortner/src/mocks"
	"github.com/kalpit-sharma-dev/url-shortner/src/model"
)

func TestCreateUrl(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockFileRepository(mockCtrl)
	urlService := NewUrlService(mockRepo)

	type args struct {
		ctx    context.Context
		urlReq model.Request
	}
	tests := []struct {
		name        string
		args        args
		setupMocks  func()
		wantUrlResp model.Response
		wantErr     bool
	}{

		{
			name: "Failure",
			args: args{
				ctx:    context.Background(),
				urlReq: model.Request{Url: ""},
			},
			setupMocks: func() {
				mockRepo.EXPECT().GetUrl(gomock.Any(), gomock.Any()).Return(model.Response{}, errors.New("Some error")).AnyTimes()
				mockRepo.EXPECT().CreateUrl(gomock.Any(), gomock.Any(), gomock.Any()).Return(model.Response{}, errors.New("Some error")).AnyTimes()

			},
			wantUrlResp: model.Response{},
			wantErr:     true,
		},
		{
			name: "Success",
			args: args{
				ctx:    context.Background(),
				urlReq: model.Request{Url: "google.com"},
			},
			setupMocks: func() {
				mockRepo.EXPECT().GetUrl(gomock.Any(), gomock.Any()).Return(model.Response{Url: "google.com", TinyUrl: "qwerty"}, nil).AnyTimes()
				mockRepo.EXPECT().CreateUrl(gomock.Any(), gomock.Any(), gomock.Any()).Return(model.Response{Url: "google.com", TinyUrl: "qwerty"}, nil).AnyTimes()

			},
			wantUrlResp: model.Response{Url: "google.com", TinyUrl: "qwerty"},
			wantErr:     false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			gotUrlResp, err := urlService.CreateUrl(tt.args.ctx, tt.args.urlReq)
			if (err != nil) != tt.wantErr {
				if err != errors.New("Some error") {
					t.Errorf("CreateUrl() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(gotUrlResp, tt.wantUrlResp) {
				t.Errorf("CreateUrl() got = %v, want %v", gotUrlResp, tt.wantUrlResp)
			}

		})
	}
}
