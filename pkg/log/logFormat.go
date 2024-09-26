package log

import "pkg/errors"

type LogFormat string

const (
	JSONFormat LogFormat = "json"
	TextFormat LogFormat = "text"
)

func (l LogFormat) Validate() error {
	switch l {
	case JSONFormat, TextFormat:
		return nil
	default:
		return errors.InternalServer.New("invalid jsonLog format")
	}
}
