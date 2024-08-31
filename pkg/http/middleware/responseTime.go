package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"server/pkg/errors"
	"server/pkg/http/chain"
	"server/pkg/http/responseWriter"
	"server/pkg/log"
)

var responseTimeMetric *prometheus.HistogramVec

func Init(serviceName string) error {

	responseTimeMetric = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   serviceName,
		Subsystem:   "",
		Name:        "http_response_time_seconds",
		Help:        "Histogram of response time in seconds.",
		ConstLabels: nil,
		Buckets:     nil,
	}, []string{"path", "status_code"})

	if err := prometheus.Register(responseTimeMetric); err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}

func ResponseTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		// Проверяем, что TimeMetric инициализирован
		if responseTimeMetric == nil {
			chain.DefaultErrorEncoder(ctx, w, errors.InternalServer.New("ResponseTime prometheus metric not initialized"))
			return
		}

		// Оборачиваем Writer в кастомный для получения статус-кода
		writerWithStatus := &responseWriter.ResponseWriterWithStatus{
			ResponseWriter: w,
			Status:         nil,
		}
		path := r.RequestURI

		start := time.Now()

		// Сохраняем статус ответа после выполнения следующего в стеке хандлера
		next.ServeHTTP(writerWithStatus, r)

		// Вычисляем продолжительность
		duration := time.Since(start)

		if writerWithStatus.Status == nil {
			log.Error(ctx, "Cannot get status code from handler")
			return
		}

		// Записываем информацию о времени ответа с использованием прометеуса
		responseTimeMetric.WithLabelValues(
			preparePath(path),
			fmt.Sprintf("%d", *writerWithStatus.Status),
		).Observe(duration.Seconds())
	})
}

// preparePath заменяет пути /rtb/ssp на rtb_ssp
func preparePath(path string) string {
	path = strings.TrimPrefix(path, "/")
	return strings.ReplaceAll(path, "/", "_")
}
