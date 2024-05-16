package applicationType

import (
	"server/app/pkg/errors"
)

type Type string

// enums:"ios,android,web,server"
const (
	IOs     = Type("ios")
	Android = Type("android")
	Web     = Type("web")
	Server  = Type("server")
)

func (t *Type) Validate() error {
	if t == nil {
		return nil
	}
	switch *t {
	case IOs, Android, Web, Server:
	default:
		return errors.BadRequest.New("Unknown application type", []errors.Option{
			errors.PathDepthOption(errors.SecondPathDepth),
			errors.ParamsOption("type", *t),
			errors.HumanTextOption("Неизвестный тип приложения"),
		}...)
	}
	return nil
}
