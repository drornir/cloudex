package htmlserver

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/drornir/cloudex/pkg/app"
	"github.com/drornir/cloudex/pkg/component"
)

func New(appl *app.App) *echo.Echo {
	e := echo.New()

	e.Use(echo.WrapMiddleware(authMiddleware))

	e.GET("/", homepage(appl))

	return e
}

func homepage(appl *app.App) func(c echo.Context) error {
	return func(c echo.Context) error {
		in := component.DocumentInput{
			Title:        "Home",
			PageNotFound: false,
		}

		comp := component.Document(in)
		b, err := component.Render(c.Request().Context(), comp)
		if err != nil {
			return fmt.Errorf("rendering homepage: %w", err)
		}

		c.HTMLBlob(http.StatusOK, b)
		return nil
	}
}
