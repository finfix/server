package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"pkg/errors"
)

func DefaultResponseEncoder(_ context.Context, w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/jsonapi; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return errors.InternalServer.Wrap(err)
	}
	return nil
}
