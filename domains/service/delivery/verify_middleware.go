package delivery

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lancer-kit/armory/api/render"
)

// VerifySomethingMiddleware is an example of custom middleware which checks parameter value from url
func VerifySomethingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mId := chi.URLParam(r, "mId")
			if mId != "test" {
				render.BadRequest(w, "Wrong param")
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), "some_param", mId))
			next.ServeHTTP(w, r)
		})
	}
}
