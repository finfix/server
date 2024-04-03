package auth

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"server/pkg/errors"
)

type MyCustomClaims struct {
	DeviceID string `json:"deviceID"`
	jwt.StandardClaims
}

func NewJWT(userId uint32, signingKey string, deviceID string, ttl time.Duration) (string, error) {

	claims := MyCustomClaims{
		DeviceID: deviceID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Subject:   strconv.Itoa(int(userId)),
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
			return nil, errors.InternalServer.NewCtx("Unexpected signing method", "%v", token.Header["alg"])
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
		return 0, "", errors.Unauthorized.WrapCtx(err, "ID: %v", idStr)
	}

	return uint32(id), claims.DeviceID, nil

}

func NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", errors.InternalServer.Wrap(err)
	}

	return fmt.Sprintf("%x", b), nil
}
