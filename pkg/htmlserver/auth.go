package htmlserver

import (
	"net/http"

	"github.com/drornir/cloudex/pkg/app"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = app.ContextWithUser(ctx, app.User{
			ID:    1,
			Email: "me@drornir.dev",
		})

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
