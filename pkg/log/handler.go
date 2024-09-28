package log

import "context"

// Handler - это интерфейс обработчика журналов.
type Handler interface {
	handle(ctx context.Context, level LogLevel, log any, opts ...Option)
}
