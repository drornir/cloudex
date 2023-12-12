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
