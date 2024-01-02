package htmlserver

import (
	"github.com/labstack/echo/v4"

	"github.com/drornir/cloudex/pkg/app"
	"github.com/drornir/cloudex/pkg/component"
)

var mockProducts = []component.Product{
	{
		Name:             "Example",
		Description:      "An example product to try things with",
		LinkToBuyLicense: "/buy-product?name=Example",
	},
}

func homepage(appl *app.App) func(c echo.Context) error {
	return func(c echo.Context) error {
		in := component.DocumentInput{
			Title:        "Home",
			PageNotFound: false,
			Content: component.MainContentInput{
				HomepageContentInput: &component.HomepageContentInput{
					Products: mockProducts,
				},
			},
		}

		comp := component.Document(in)
		return renderComp(c, comp)
	}
}
