package jwtManager

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"server/app/pkg/errors"
)

type JWTManager struct {
	accessTokenSigningKey []byte
	ttls                  map[TokenType]time.Duration
}

var jwtManager *JWTManager

func Init(
	accessTokenSigningKey []byte,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) {
	jwtManager = &JWTManager{
		accessTokenSigningKey: accessTokenSigningKey,
		ttls: map[TokenType]time.Duration{
			RefreshToken: refreshTokenTTL,
			AccessToken:  accessTokenTTL,
		},
	}
}

type MyCustomClaims struct {
	DeviceID string `json:"deviceID"`
	jwt.StandardClaims
}

type TokenType int

const (
	RefreshToken = iota + 1
	AccessToken
)

func NewJWT(tokenType TokenType, userID uint32, deviceID string) (string, error) {

	if jwtManager == nil {
		return "", errors.InternalServer.New("JWTManager is not initialized")
	}

	claims := MyCustomClaims{
		DeviceID: deviceID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(jwtManager.ttls[tokenType]).Unix(),
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

func Parse(reqToken string) (uint32, string, error) {

	if jwtManager == nil {
		return 0, "", errors.InternalServer.New("JWTManager is not initialized", []errors.Option{
			errors.PathDepthOption(errors.SecondPathDepth),
		}...)
	}

	if reqToken == "" {
		return 0, "", errors.Unauthorized.New("JWT-token is empty", []errors.Option{
			errors.PathDepthOption(errors.SecondPathDepth),
		}...)
	}

	token, jwtErr := jwt.ParseWithClaims(reqToken, &MyCustomClaims{}, func(token *jwt.Token) (i any, err error) { //nolint:exhaustruct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.InternalServer.New("Unexpected signing method", []errors.Option{
				errors.ParamsOption("token", token.Header["alg"]),
				errors.PathDepthOption(errors.SecondPathDepth),
			}...)
		}

		return jwtManager.accessTokenSigningKey, nil
	})
	if jwtErr != nil {
		if !errors.As(jwtErr, jwt.ValidationErrorExpired) {
			return 0, "", errors.BadRequest.Wrap(jwtErr, []errors.Option{
				errors.PathDepthOption(errors.SecondPathDepth),
			}...)
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
