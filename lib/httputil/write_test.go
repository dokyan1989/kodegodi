package httputil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	now := GetTimeNow()

	type args struct {
		w    *httptest.ResponseRecorder
		v    SampleResponse
		code int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "status ok",
			args: args{
				w: httptest.NewRecorder(),
				v: SampleResponse{
					Code:    0,
					Message: "success",
					Data: SampleResponseData{
						FirstName: "firstname",
						LastName:  "lastname",
						CreatedAt: now,
					},
				},
				code: http.StatusOK,
			},
		},
		{
			name: "status bad request",
			args: args{
				w: httptest.NewRecorder(),
				v: SampleResponse{
					Code:    1,
					Message: "invalid parameters",
					Data:    SampleResponseData{},
				},
				code: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteJSON(tt.args.w, tt.args.v, tt.args.code)

			res := tt.args.w.Result()
			defer res.Body.Close()

			var actual SampleResponse
			json.NewDecoder(res.Body).Decode(&actual)

			if !reflect.DeepEqual(tt.args.v, actual) {
				t.Errorf("WriteJSON() | v = %v, actual %v", tt.args.v, actual)
				return
			}
			if tt.args.code != res.StatusCode {
				t.Errorf("WriteJSON() | code = %v, actual %v", tt.args.code, res.StatusCode)
			}
		})
	}
}
