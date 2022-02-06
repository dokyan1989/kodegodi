package httputil

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadBodyJSON(resp *http.Response, v interface{}) error {
	body, err := ReadBodyBytes(resp)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

func ReadBodyString(resp *http.Response) (string, error) {
	body, err := ReadBodyBytes(resp)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ReadBodyBytes(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return []byte{}, err
	}

	return body, err
}
