package smfp

import "errors"

var (
	ErrEmptyPipe         = errors.New("smfp: a request has been made to an empty pipe")
	ErrInvalidOutputType = errors.New("smfp: output type is invalid")
)
