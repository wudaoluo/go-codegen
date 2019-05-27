package internal

import (
	"github.com/pkg/errors"
)

var ERROR_TPL_NOT_FOUND = errors.New("template does not exist")

type TPlError struct {
	TplName string
	Op      string
	Err     error
}

func (e *TPlError) Error() string {
	if e == nil {
		return "<nil>"
	}

	s := e.Op
	if e.TplName != "" {
		s += " " + e.TplName
	}

	s += ": " + e.Err.Error()

	return s
}
