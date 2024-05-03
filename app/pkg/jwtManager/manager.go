package jwtManager

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"server/app/config"
	"server/app/pkg/errors"
	"server/app/pkg/hasher"
)

type JWTManager struct {
	signingKey []byte
	ttl        time.Duration
}

var jwtManager JWTManager

func Init(signingKey []byte, ttl time.Duration) {
	jwtManager = JWTManager{
		signingKey: signingKey,
		ttl:        ttl,
	}
}

type MyCustomClaims struct {
	DeviceID string `json:"deviceID"`
	jwt.StandardClaims
}

func NewJWT(userID uint32, deviceID string) (string, error) {

	claims := MyCustomClaims{
		DeviceID: deviceID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtManager.ttl).Unix(),
			Subject:   strconv.Itoa(int(userID)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(jwtManager.signingKey)
	if err != nil {
		return "", errors.InternalServer.Wrap(err)
	}

	return tokenStr, nil
}

func Parse(accessToken string) (uint32, string, error) {

	if accessToken == "" {
		return 0, "", errors.Unauthorized.New("JWT-token is empty")
	}

	token, jwtErr := jwt.ParseWithClaims(accessToken, &MyCustomClaims{}, func(token *jwt.Token) (i any, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.InternalServer.New("Unexpected signing method", errors.Options{
				Params: map[string]any{"token": token.Header["alg"]},
			})
		}

		return jwtManager.signingKey, nil
	})
	if jwtErr != nil {
		if !errors.As(jwtErr, jwt.ValidationErrorExpired) {
			return 0, "", errors.BadRequest.Wrap(jwtErr)
		} else {
			jwtErr = errors.Unauthorized.Wrap(jwtErr)
		}
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return 0, "", errors.InternalServer.New("Error get user claims from token")
	}

	idStr := claims.Subject
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, "", errors.BadRequest.Wrap(err, errors.Options{
			Params: map[string]any{"ID": idStr},
		})
	}

	return uint32(id), claims.DeviceID, jwtErr
}

func NewRefreshToken() (string, time.Time, error) {

	const refreshTokenLength = 64

	token, err := hasher.GenerateRandomBytes(refreshTokenLength)
	if err != nil {
		return "", time.Now(), err
	}

	// Получаем время жизни refresh token
	refreshDur, err := time.ParseDuration(config.GetConfig().Token.RefreshTokenTTL)
	if err != nil {
		return "", time.Now(), err
	}
	refreshTokenExpiresAt := time.Now().Add(refreshDur)

	return string(token), refreshTokenExpiresAt, nil
}
