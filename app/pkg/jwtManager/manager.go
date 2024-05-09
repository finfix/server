package jwtManager

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"server/app/pkg/errors"
	"server/app/pkg/hasher"
)

type JWTManager struct {
	accessTokenSigningKey []byte
	accessTokenTTL        time.Duration
	refreshTokenTTL       time.Duration
}

var jwtManager *JWTManager

func Init(
	accessTokenSigningKey []byte,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) {
	jwtManager = &JWTManager{
		accessTokenSigningKey: accessTokenSigningKey,
		accessTokenTTL:        accessTokenTTL,
		refreshTokenTTL:       refreshTokenTTL,
	}
}

type MyCustomClaims struct {
	DeviceID string `json:"deviceID"`
	jwt.StandardClaims
}

func NewJWT(userID uint32, deviceID string) (string, error) {

	if jwtManager == nil {
		return "", errors.InternalServer.New("JWTManager is not initialized")
	}

	claims := MyCustomClaims{
		DeviceID: deviceID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(jwtManager.accessTokenTTL).Unix(),
			Id:        "",
			IssuedAt:  0,
			Issuer:    "",
			NotBefore: 0,
			Subject:   strconv.Itoa(int(userID)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(jwtManager.accessTokenSigningKey)
	if err != nil {
		return "", errors.InternalServer.Wrap(err)
	}

	return tokenStr, nil
}

func Parse(accessToken string) (uint32, string, error) {

	if jwtManager == nil {
		return 0, "", errors.InternalServer.New("JWTManager is not initialized")
	}

	if accessToken == "" {
		return 0, "", errors.Unauthorized.New("JWT-token is empty")
	}

	token, jwtErr := jwt.ParseWithClaims(accessToken, &MyCustomClaims{}, func(token *jwt.Token) (i any, err error) { //nolint:exhaustruct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.InternalServer.New("Unexpected signing method", []errors.Option{
				errors.ParamsOption("token", token.Header["alg"]),
			}...)
		}

		return jwtManager.accessTokenSigningKey, nil
	})
	if jwtErr != nil {
		if !errors.As(jwtErr, jwt.ValidationErrorExpired) {
			return 0, "", errors.BadRequest.Wrap(jwtErr)
		} else {
			jwtErr = errors.Unauthorized.Wrap(jwtErr, []errors.Option{
				errors.PathDepthOption(errors.SecondPathDepth),
			}...)
		}
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return 0, "", errors.InternalServer.New("Error get user claims from token")
	}

	idStr := claims.Subject
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, "", errors.BadRequest.Wrap(err, []errors.Option{
			errors.ParamsOption("ID", idStr),
		}...)
	}

	return uint32(id), claims.DeviceID, jwtErr
}

func NewRefreshToken() (string, time.Time, error) {

	if jwtManager == nil {
		return "", time.Now(), errors.InternalServer.New("JWTManager is not initialized")
	}

	const refreshTokenLength = 64

	bytes, err := hasher.GenerateRandomBytes(refreshTokenLength)
	if err != nil {
		return "", time.Now(), err
	}

	// Получаем время жизни refresh token
	refreshDur, err := time.ParseDuration(jwtManager.refreshTokenTTL.String())
	if err != nil {
		return "", time.Now(), err
	}
	refreshTokenExpiresAt := time.Now().Add(refreshDur)

	return fmt.Sprintf("%x", bytes[:refreshTokenLength]), refreshTokenExpiresAt, nil
}
