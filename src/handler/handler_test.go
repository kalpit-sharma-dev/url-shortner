package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kalpit-sharma-dev/url-shortner/src/mocks"
	"github.com/kalpit-sharma-dev/url-shortner/src/model"
)

func TestHandleCreateUrl(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUrlService(ctrl)
	urlHandler := NewHandler(mockRepo)

	var successReq = model.Request{
		Url: "google.com",
	}

	tests := []struct {
		name           string
		req            model.Request
		setupMocks     func()
		expectedStatus int
	}{
		{
			name: "Success",
			req:  successReq,
			setupMocks: func() {
				mockRepo.EXPECT().CreateUrl(gomock.Any(), gomock.Any()).Return(model.Response{}, nil).AnyTimes()
			},
			expectedStatus: 201,
		},
		// {
		// 	name: "Error from service",
		// 	req:  successReq,
		// 	setupMocks: func() {
		// 		mockRepo.EXPECT().CreateUrl(gomock.Any(), gomock.Any()).Return(model.Response{}, errors.New("Some error from service")).AnyTimes()
		// 	},
		// 	expectedStatus: 500,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			b, _ := json.Marshal(tt.req)
			rr := httptest.NewRecorder()
			httpReq, _ := http.NewRequest("POST", "/url", bytes.NewBuffer(b))
			urlHandler.HandleCreateUrl(rr, httpReq)
			if tt.expectedStatus != rr.Code {
				t.Errorf("Error got = %v and want = %v", rr.Code, tt.expectedStatus)
			}
		})
	}
}
