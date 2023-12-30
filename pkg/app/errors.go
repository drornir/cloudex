package app

import (
	"context"
	"fmt"
)

type ErrorUserIsMissing struct {
	Context context.Context
}

func (e ErrorUserIsMissing) Error() string {
	return fmt.Sprintf("%T", e)
}

type ErrorSQL struct {
	Msg string
	Err error
}

func (e ErrorSQL) Error() string {
	return fmt.Sprintf("sql: %s", e.Err)
}
func (e ErrorSQL) Unwrap() error {
	return e.Err
}
