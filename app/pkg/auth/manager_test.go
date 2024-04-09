package auth

import (
	"fmt"
	"testing"

	"server/app/pkg/testingFunc"
)

func TestNewRefreshToken(t *testing.T) {

	mapOfTokens := make(map[string]struct{})

	count := 1000

	for i := 1; i <= count; i++ {
		t.Run(fmt.Sprintf("%v. Создание уникального токена", i), func(t *testing.T) {
			token, err := NewRefreshToken()
			if testingFunc.CheckError(t, nil, err) {
				return
			}
			if _, ok := mapOfTokens[token]; ok {
				t.Errorf("Сгенерированный токен не уникален. Токен: %v", token)
			}
			mapOfTokens[token] = struct{}{}
		})
	}
}
