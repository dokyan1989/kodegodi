package httputil

import (
	"errors"
	"fmt"
)

var (
	ErrResponseFailed = func(statusCode int) error {
		return fmt.Errorf("response failed with status code: %d", statusCode)
	}

	ErrUnsupportedEncodingType = errors.New("encoding type is not supported")
)
