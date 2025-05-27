package repository

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rieshbissessur/dependency-mock-generator/internal/model"
	"github.com/rieshbissessur/dependency-mock-generator/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestAddGetMappingModel(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		path       string
		statusCode int64
		response   map[string]any
		wantErr    bool
		mockErr    error
	}{
		{
			name:       "successful stub",
			url:        "http://localhost:8080",
			path:       "/test",
			statusCode: 200,
			response:   map[string]any{"key": "value"},
			wantErr:    false,
			mockErr:    nil,
		},
		{
			name:       "stub error",
			url:        "http://localhost:8080",
			path:       "/test",
			statusCode: 200,
			response:   map[string]any{"key": "value"},
			wantErr:    true,
			mockErr:    errors.New("stub error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mock.NewMockWiremockClient(ctrl)
			mockClient.EXPECT().
				StubFor(gomock.Any()).
				Return(tt.mockErr)

			origNewClient := NewWiremockClient
			NewWiremockClient = func(url string) WiremockClient {
				return mockClient
			}
			defer func() { NewWiremockClient = origNewClient }()

			err := AddGetMappingModel(tt.url, tt.path, tt.statusCode, tt.response)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAddPostMappingModel(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		path       string
		statusCode int64
		response   map[string]any
		wantErr    bool
		mockErr    error
	}{
		{
			name:       "successful stub",
			url:        "http://localhost:8080",
			path:       "/test",
			statusCode: 200,
			response:   map[string]any{"key": "value"},
			wantErr:    false,
			mockErr:    nil,
		},
		{
			name:       "stub error",
			url:        "http://localhost:8080",
			path:       "/test",
			statusCode: 200,
			response:   map[string]any{"key": "value"},
			wantErr:    true,
			mockErr:    errors.New("stub error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mock.NewMockWiremockClient(ctrl)
			mockClient.EXPECT().
				StubFor(gomock.Any()).
				Return(tt.mockErr)

			origNewClient := NewWiremockClient
			NewWiremockClient = func(url string) WiremockClient {
				return mockClient
			}
			defer func() { NewWiremockClient = origNewClient }()

			err := AddPostMappingModel(tt.url, tt.path, tt.statusCode, tt.response)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestClearAllMappingModels(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
		mockErr error
	}{
		{
			name:    "successful clear",
			url:     "http://localhost:8080",
			wantErr: false,
			mockErr: nil,
		},
		{
			name:    "clear error",
			url:     "http://localhost:8080",
			wantErr: true,
			mockErr: errors.New("clear error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mock.NewMockWiremockClient(ctrl)
			mockClient.EXPECT().
				Reset().
				Return(tt.mockErr)

			origNewClient := NewWiremockClient
			NewWiremockClient = func(url string) WiremockClient {
				return mockClient
			}
			defer func() { NewWiremockClient = origNewClient }()

			err := ClearAllMappingModels(tt.url)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAllMappingModels(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		wantErr    bool
		mockResp   *http.Response
		mockErr    error
		wantResult string
	}{
		{
			name:    "successful get",
			url:     "http://localhost:8080",
			wantErr: false,
			mockResp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"mappings": []}`)),
			},
			mockErr:    nil,
			wantResult: `{"mappings": []}`,
		},
		{
			name:     "http error",
			url:      "http://localhost:8080",
			wantErr:  true,
			mockResp: nil,
			mockErr:  errors.New("http error"),
		},
		{
			name: "non-200 status",
			url:  "http://localhost:8080",
			mockResp: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader("")),
			},
			wantErr: true,
			mockErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockHTTP := mock.NewMockHTTPClient(ctrl)
			mockHTTP.EXPECT().
				Get(tt.url+"/__admin/mappings").
				Return(tt.mockResp, tt.mockErr)

			origHTTP := HttpClient
			HttpClient = mockHTTP
			defer func() { HttpClient = origHTTP }()

			result, err := GetAllMappingModels(tt.url)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantResult, result)
			}
		})
	}
}

func TestImportMappingModels(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		mapping  model.Mapping
		wantErr  bool
		mockResp *http.Response
		mockErr  error
	}{
		{
			name: "successful import",
			url:  "http://localhost:8080",
			mapping: model.Mapping{
				ID: "test-id",
				Request: model.Request{
					UrlPattern: "/test",
					Method:     "GET",
				},
				Response: model.Response{
					Status:   200,
					JsonBody: map[string]string{"key": "value"},
				},
			},
			wantErr: false,
			mockResp: &http.Response{
				StatusCode: http.StatusCreated,
				Body:       io.NopCloser(strings.NewReader("")),
			},
			mockErr: nil,
		},
		{
			name: "http error",
			url:  "http://localhost:8080",
			mapping: model.Mapping{
				ID: "test-id",
			},
			wantErr:  true,
			mockResp: nil,
			mockErr:  errors.New("http error"),
		},
		{
			name: "non-201 status",
			url:  "http://localhost:8080",
			mapping: model.Mapping{
				ID: "test-id",
			},
			mockResp: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader("")),
			},
			wantErr: true,
			mockErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockHTTP := mock.NewMockHTTPClient(ctrl)
			mockHTTP.EXPECT().
				Post(tt.url+"/__admin/mappings", "application/json", gomock.Any()).
				Return(tt.mockResp, tt.mockErr)

			origHTTP := HttpClient
			HttpClient = mockHTTP
			defer func() { HttpClient = origHTTP }()

			err := ImportMappingModels(tt.url, tt.mapping)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
