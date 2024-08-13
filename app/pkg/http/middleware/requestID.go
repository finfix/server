package middleware

import (
	"net/http"

	"github.com/google/uuid"

	"server/app/pkg/contextKeys"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Генерируем UUID
		uuid := uuid.New().String()

		// Добавляем UUID в заголовок ответа
		w.Header().Set("X-Request-ID", uuid)

		// Добавляем UUID в контекст
		ctx := contextKeys.SetRequestID(r.Context(), uuid)

		// Вызываем следующий хэндлер
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
