package htmlserver

import (
	"net/http"

	"github.com/a-h/templ"
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

	e.GET("/buy-product", buyProductPage(appl))
	e.POST("/buy-product", createNewLicense(appl))

	e.StaticFS("/assets", appl.AssetsFS)

	return e
}

func renderComp(c echo.Context, comp templ.Component) error {
	b, err := component.Render(c.Request().Context(), comp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "rendering html").
			SetInternal(err)
	}

	c.HTMLBlob(http.StatusOK, b)
	return nil
}
