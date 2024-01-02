package htmlserver

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/drornir/cloudex/pkg/app"
	"github.com/drornir/cloudex/pkg/component"
	"github.com/drornir/cloudex/pkg/product"
)

func ProductToComponent(prod product.Product) component.Product {
	return component.Product{
		Name:             prod.Name(),
		Description:      prod.Description(),
		LinkToBuyLicense: fmt.Sprintf("/buy-product?name=%s", prod.Name()),
	}
}

func buyProductPage(appl *app.App) func(c echo.Context) error {
	return func(c echo.Context) error {
		prodName := c.QueryParam("name")
		if prodName == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "query param 'name' for product name is required")
		}

		prod, ok := product.Products()[prodName]
		if !ok {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("product with name %q wasn't found", prodName))
		}

		in := component.DocumentInput{
			Title:        "Buy",
			PageNotFound: false,
			Content: component.MainContentInput{
				BuyProductContentInput: &component.BuyProductContentInput{
					Product: ProductToComponent(prod),
				},
			},
		}

		comp := component.Document(in)
		return renderComp(c, comp)
	}
}

func createNewLicense(appl *app.App) func(c echo.Context) error {
	return func(c echo.Context) error {
		var frm struct {
			Name string `form:"name"`
		}
		if err := c.Bind(&frm); err != nil {
			return err
		}

		if frm.Name == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "query param 'name' for product name is required")
		}

		l, err := appl.BuyProduct(c.Request().Context(), frm.Name)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "buying license").
				SetInternal(err)
		}

		comp := component.BuyProductResponse(l)
		return renderComp(c, comp)
	}
}
