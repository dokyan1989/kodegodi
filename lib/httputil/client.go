package httputil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Get(url string) (*http.Response, error) {
	return http.Get(url)
}

func Post(url, contentType string, v interface{}) (*http.Response, error) {
	body, err := getBodyReader(contentType, v)
	if err != nil {
		return nil, err
	}
	return http.Post(url, contentType, body)
}

func Put(url, contentType string, v interface{}) (*http.Response, error) {
	body, err := getBodyReader(contentType, v)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(HeaderContentType, contentType)
	return client.Do(req)
}

func Patch(url, contentType string, v interface{}) (*http.Response, error) {
	body, err := getBodyReader(contentType, v)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(HeaderContentType, contentType)
	return client.Do(req)
}

func Delete(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func getBodyReader(contentType string, v interface{}) (io.Reader, error) {
	var buf bytes.Buffer

	switch contentType {
	case MIMETypeApplicationJSON:
		err := json.NewEncoder(&buf).Encode(v)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrUnsupportedEncodingType
	}

	return &buf, nil
}
