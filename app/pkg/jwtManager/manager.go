package jwtManager

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

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
	UserID   uint32 `json:"userID"`
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
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(jwtManager.ttls[tokenType]).Unix(),
			Id:        uuid.New().String(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "bonavii.com",
			NotBefore: 0,
			Subject:   "",
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
		return 0, "", errors.InternalServer.New("JWTManager is not initialized",
			errors.StackTraceOption(errors.PreviousCaller),
		)
	}

	if reqToken == "" {
		return 0, "", errors.Unauthorized.New("JWT-token is empty",
			errors.StackTraceOption(errors.PreviousCaller),
		)
	}

	token, jwtErr := jwt.ParseWithClaims(reqToken, &MyCustomClaims{}, func(token *jwt.Token) (i any, err error) { //nolint:exhaustruct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.InternalServer.New("Unexpected signing method",
				errors.ParamsOption("token", token.Header["alg"]),
				errors.StackTraceOption(errors.PreviousCaller),
			)
		}

		return jwtManager.accessTokenSigningKey, nil
	})
	if jwtErr != nil {
		if !errors.As(jwtErr, jwt.ValidationErrorExpired) {
			return 0, "", errors.BadRequest.Wrap(jwtErr,
				errors.StackTraceOption(errors.PreviousCaller),
			)
		} else {
			jwtErr = errors.Unauthorized.Wrap(jwtErr,
				errors.StackTraceOption(errors.PreviousCaller),
			)
		}
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return 0, "", errors.InternalServer.New("Error get user claims from token")
	}

	return claims.UserID, claims.DeviceID, jwtErr
}
