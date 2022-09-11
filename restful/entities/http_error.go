package entities

import "errors"

var (
	InternalError     = errors.New("Internal Server Error")
	ParamInvalid      = errors.New("Params Is Invalid")
	MethodNotAllowErr = errors.New("Method Not Allow")
)
