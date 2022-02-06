package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func (c Client) UpperCase(word string) (string, error) {
	res, err := http.Get(c.url + "/upper?word=" + word)
	if err != nil {
		return "", fmt.Errorf("unable to complete Get request, error = %s", err.Error())
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read response data, error = %s", err.Error())
	}

	return string(out), nil
}
