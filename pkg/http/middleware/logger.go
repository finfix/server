package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"pkg/contextKeys"
	"pkg/log"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		var requestID string
		if _requestID := contextKeys.GetRequestID(ctx); _requestID != nil {
			requestID = *_requestID
		}

		log.Info(ctx, fmt.Sprintf("%s [%s] %s", r.URL.Path, strings.ToLower(r.Method), requestID))

		next.ServeHTTP(w, r)
	})
}
