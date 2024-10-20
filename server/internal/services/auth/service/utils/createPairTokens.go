package utils

import (
	"context"

	"pkg/jwtManager"

	authModel "server/internal/services/auth/model"
)

func CreatePairTokens(_ context.Context, userID uint32, deviceID string) (tokens authModel.Tokens, err error) {

	// Создаем Access token
	tokens.AccessToken, err = jwtManager.GenerateToken(jwtManager.AccessToken, userID, deviceID)
	if err != nil {
		return tokens, err
	}

	// Создаем refresh token
	tokens.RefreshToken, err = jwtManager.GenerateToken(jwtManager.RefreshToken, userID, deviceID)
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}
