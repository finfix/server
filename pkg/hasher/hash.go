package hasher

import (
	"crypto/sha256"
	"fmt"

	"server/pkg/errors"
)

func Hash(password, shasalt string) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", errors.InternalServer.Wrap(err)
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(shasalt))), nil
}
