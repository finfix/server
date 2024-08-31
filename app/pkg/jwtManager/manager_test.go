package jwtManager

import (
	"fmt"
	"testing"
	"time"

	"server/app/pkg/testUtils"
)

func TestNewJWTToken(t *testing.T) {

	Init([]byte(""), time.Hour, time.Hour)

	count := 1000
	mapOfTokens := make(map[string]struct{}, count)

	for i := 1; i <= count; i++ {
		t.Run(fmt.Sprintf("%v. Создание уникального токена", i), func(t *testing.T) {
			token, err := NewJWT(RefreshToken, 1, "deviceID")
			if testUtils.CheckError(t, nil, err, false) {
				return
			}
			if _, ok := mapOfTokens[token]; ok {
				t.Errorf("Сгенерированный токен не уникален. Токен: %v", token)
			}
			mapOfTokens[token] = struct{}{}
		})
	}
}
