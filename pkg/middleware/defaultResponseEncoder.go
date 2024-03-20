package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"server/pkg/errors"
)

func DefaultResponseEncoder(_ context.Context, w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return errors.InternalServer.Wrap(err)
	}
	return nil
}
