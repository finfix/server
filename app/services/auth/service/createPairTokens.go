package service

import (
	"context"

	"server/app/pkg/jwtManager"
	authModel "server/app/services/auth/model"
)

func (s *Service) createPairTokens(_ context.Context, userID uint32, deviceID string) (tokens authModel.Tokens, err error) {

	// Создаем Access token
	tokens.AccessToken, err = jwtManager.NewJWT(jwtManager.AccessToken, userID, deviceID)
	if err != nil {
		return tokens, err
	}

	// Создаем refresh token
	tokens.RefreshToken, err = jwtManager.NewJWT(jwtManager.RefreshToken, userID, deviceID)
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}
