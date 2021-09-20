package httputil

import (
	"example.com/grpcdemo/config"
	"example.com/grpcdemo/utils/dlog"
	"testing"
	"time"
)

func init() {
	config.Mode = config.ModeQa
	dlog.SetTopic("test")
	dlog.DebugLog(true) // 开启debug日志
}

func TestHttpGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name        string
		args        args
		wantCode    int
		wantErr     bool
	}{
		{
			name: "test_http_get",
			args: args{url: "https://www.baidu.com"},
			wantCode: 200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotContent, err := HttpGet(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCode != tt.wantCode {
				t.Errorf("HttpGet() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
			dlog.Info("content", gotContent)
		})
	}
}

func TestHttpGetWithOption(t *testing.T) {
	type args struct {
		url     string
		options []*Option
	}
	tests := []struct {
		name        string
		args        args
		wantCode    int
		wantErr     bool
	}{
		{
			name: "test_http_get_option",
			args: args{
				url:     "https://www.baidu.com",
				options: []*Option{WithTimeout(1 * time.Second)},
			},
			wantCode: 200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotContent, err := HttpGet(tt.args.url, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpGetWithOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCode != tt.wantCode {
				t.Errorf("HttpGetWithOption() gotCode = %v, want %v", gotCode, tt.wantCode)
			}

			dlog.Info(gotCode, gotContent)
		})
	}
}

func TestHttpJsonPost(t *testing.T) {
	type args struct {
		url     string
		data    string
		timeout time.Duration
	}
	tests := []struct {
		name        string
		args        args
		wantCode    int
		wantContent string
		wantErr     bool
	}{
		{
			name: "test_http_get_option",
			args: args{
				url:     "https://www.baidu.com",
				data: "123",
				timeout: 1 * time.Second,
			},
			wantCode: 200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotContent, err := HttpJsonPost(tt.args.url, tt.args.data, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpJsonPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCode != tt.wantCode {
				t.Errorf("HttpJsonPost() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
			dlog.Info("111111111111111", gotCode, gotContent)
		})
	}
}