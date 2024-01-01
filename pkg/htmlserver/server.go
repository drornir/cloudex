package htmlserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/drornir/cloudex/pkg/app"
	"github.com/drornir/cloudex/pkg/component"
)

func New(appl *app.App) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(authMiddleware(appl))

	e.GET("/", homepage(appl))

	e.StaticFS("/assets", appl.AssetsFS)

	return e
}

func homepage(appl *app.App) func(c echo.Context) error {
	return func(c echo.Context) error {
		in := component.DocumentInput{
			Title:        "Home",
			PageNotFound: false,
			Content: component.MainContentInput{
				Products: []component.Product{
					{
						Name:             "Example",
						Description:      "An example product to try things with",
						LinkToBuyLicense: "/buy-product?name=Example",
					},
				},
			},
		}

		comp := component.Document(in)
		b, err := component.Render(c.Request().Context(), comp)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "rendering homepage").
				SetInternal(err)
		}

		c.HTMLBlob(http.StatusOK, b)
		return nil
	}
}
