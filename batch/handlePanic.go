package batch

import (
	"errors"
	"fmt"
)

type handlePanic struct {
}

func (h *handlePanic) handlePanicError(err *error) {
	if e := recover(); e != nil {
		if v, ok := e.(error); ok {
			*err = v
		} else {
			*err = errors.New(fmt.Sprintf("%v", e))
		}
	}
}
