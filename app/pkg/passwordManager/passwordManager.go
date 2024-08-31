package passwordManager

import (
	"bytes"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"server/app/pkg/errors"
)

const saltSize = 16

func CreateNewPassword(password, generalSalt []byte) ([]byte, []byte, error) {

	if len(password) > 72 {
		availableLength := len(password) - len(generalSalt) - saltSize
		return nil, nil, errors.BadRequest.New(fmt.Sprintf("Длина пароля не должна превышать %v символа", availableLength))
	}

	userSalt, err := GenerateRandomBytes(saltSize)
	if err != nil {
		return nil, nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword(saltPassword(password, userSalt, generalSalt), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, errors.InternalServer.Wrap(err)
	}

	return passwordHash, userSalt, nil
}

func CompareHashAndPassword(hash, password, userSalt, generalSalt []byte) error {

	if err := bcrypt.CompareHashAndPassword(hash, saltPassword(password, userSalt, generalSalt)); err != nil {
		return errors.BadRequest.Wrap(err,
			errors.HumanTextOption("Неверно введен логин или пароль"),
		)
	}
	return nil
}

func saltPassword(password, userSalt, generalSalt []byte) []byte {
	return bytes.Join([][]byte{userSalt, password, generalSalt}, nil)
}

func GenerateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	return b, nil
}
