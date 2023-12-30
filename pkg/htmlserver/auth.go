package htmlserver

import (
	"database/sql"
	"errors"

	"github.com/drornir/cloudex/pkg/app"
	"github.com/labstack/echo/v4"
)

func authMiddleware(appl *app.App) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			userID := int64(1) // TODO

			u, err := appl.DB.GetUser(ctx, userID)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return err
				}
				return next(c)
			}

			uu := app.User(u)
			ctx = app.ContextWithUser(ctx, uu)

			r := c.Request().WithContext(ctx)
			c.SetRequest(r)
			return next(c)
		}
	}
}
