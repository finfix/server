package hasher

import (
	"crypto/sha1"
	"fmt"

	"server/pkg/errors"
)

func Hash(password, shasalt string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", errors.InternalServer.Wrap(err)
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(shasalt))), nil
}
