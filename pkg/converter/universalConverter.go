package converter

import (
	"encoding/json"
	"pkg/errors"
)

func Convert[T any](dest T, src any) (T, error) {
	bytesJson, err := json.Marshal(src)
	if err != nil {
		return dest, errors.InternalServer.Wrap(err)
	}

	err = json.Unmarshal(bytesJson, &dest)
	if err != nil {
		return dest, errors.InternalServer.Wrap(err)
	}
	return dest, nil
}
