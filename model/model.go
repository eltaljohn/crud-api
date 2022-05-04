package model

import "errors"

var (
	ErrPersonCanNotBeNil    = errors.New("the person can not be nil")
	ErrIDPersonDoesNotExist = errors.New("the person does no exist")
)
