package httputil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		svr *httptest.Server
	}
	type want struct {
		body       string
		statusCode int
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				svr: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprint(w, "Hello World")
				})),
			},
			want: want{
				body:       "Hello World",
				statusCode: http.StatusOK,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.svr.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotBody, err := ReadBodyString(got)
			if err != nil {
				t.Error("Cannot read body response")
			}

			actual := want{gotBody, got.StatusCode}
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("Get() = %v, want %v", actual, tt.want)
			}
		})
	}
}
