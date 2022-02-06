package httputil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestReadBodyJSON(t *testing.T) {
	now := GetTimeNow()

	type args struct {
		resp *http.Response
		v    SampleResponse
	}
	tests := []struct {
		name     string
		args     args
		jsonData string
		want     SampleResponse
		wantErr  bool
	}{
		{
			name:     "success",
			jsonData: fmt.Sprintf(`{"code":0,"message":"success","data":{"firstName": "firstname","lastName":"lastname","createdAt":"%s"}}`, now.Format(time.RFC3339)),
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusOK,
				},
				v: SampleResponse{},
			},
			want: SampleResponse{
				Code:    0,
				Message: "success",
				Data: SampleResponseData{
					FirstName: "firstname",
					LastName:  "lastname",
					CreatedAt: now,
				},
			},
			wantErr: false,
		},
		{
			name:     "error",
			jsonData: `{"code":1,"message":"invalid parameters","data":{}`,
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusBadRequest,
				},
				v: SampleResponse{},
			},
			want:    SampleResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.args.resp.Body = ioutil.NopCloser(bytes.NewReader([]byte(tt.jsonData)))

			if err := ReadBodyJSON(tt.args.resp, &tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ReadBodyJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// res, ok := tt.args.v.(*SampleResponse)
			// if !ok {
			// 	t.Error("Response body is not correct type")
			// }

			if !reflect.DeepEqual(tt.args.v, tt.want) {
				t.Errorf("ReadBodyJSON() = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}
