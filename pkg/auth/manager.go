package auth

import (
	"crypto/rand"
	"encoding/base32"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"server/pkg/errors"
)

type MyCustomClaims struct {
	DeviceID string `json:"deviceId"`
	jwt.StandardClaims
}

func NewJWT(userID uint32, signingKey string, deviceID string, ttl time.Duration) (string, error) {

	claims := MyCustomClaims{
		DeviceID: deviceID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Subject:   strconv.Itoa(int(userID)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", errors.InternalServer.Wrap(err)
	}

	return tokenStr, nil
}

func Parse(accessToken, signingKey string) (uint32, string, error) {

	if accessToken == "" {
		return 0, "", errors.Unauthorized.New("JWT-token is empty")
	}

	token, err := jwt.ParseWithClaims(accessToken, &MyCustomClaims{}, func(token *jwt.Token) (i any, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.InternalServer.New("Unexpected signing method", errors.Options{
				Params: map[string]any{"token": token.Header["alg"]},
			})
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", errors.Unauthorized.Wrap(err)
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return 0, "", errors.InternalServer.New("Error get user claims from token")
	}

	idStr := claims.Subject
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, "", errors.Unauthorized.Wrap(err, errors.Options{
			Params: map[string]any{"ID": idStr},
		})
	}

	return uint32(id), claims.DeviceID, nil

}

func NewRefreshToken() (string, error) {
	const (
		refreshTokenLength = 64
		countBytes         = 64
	)
	randomBytes := make([]byte, countBytes)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", errors.InternalServer.Wrap(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:refreshTokenLength], nil
}
