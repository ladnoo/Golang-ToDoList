package todo

import "errors"

var ErrTaskNotFound error = errors.New("todo not found")
var ErrTaskAlreadyExists error = errors.New("todo already exists")
