package middleware

import (
	"net/http"

	"pkg/http/chain"
	"pkg/panicRecover"
)

func PanicRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		defer panicRecover.PanicRecover(func(err error) {
			chain.DefaultErrorEncoder(ctx, w, err)
		})

		next.ServeHTTP(w, r)
	})
}
