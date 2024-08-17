package responseWriter

import "net/http"

type ResponseWriterWithStatus struct {
	http.ResponseWriter
	Status *int
}

func (r *ResponseWriterWithStatus) WriteHeader(statusCode int) {
	r.Status = &statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
