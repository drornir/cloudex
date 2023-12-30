package component

import (
	"bytes"
	"context"

	"github.com/a-h/templ"
)

func Render(ctx context.Context, comp templ.Component) ([]byte, error) {
	var b bytes.Buffer
	err := comp.Render(ctx, &b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
