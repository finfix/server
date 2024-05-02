package service

import (
	"context"
	"time"

	"server/app/pkg/errors"
	"server/app/services/auth/model"
)

// RefreshTokens обновляет токены доступа в базе данных
func (s *Service) RefreshTokens(ctx context.Context, req model.RefreshTokensReq) (newTokens model.RefreshTokensRes, err error) {

	// Получаем сессию пользователя
	session, err := s.authRepository.GetSession(ctx, req)
	if err != nil {
		return newTokens, err
	}

	// Проверяем, есть ли вообще сессия
	if session.ID == 0 {
		return newTokens, errors.Unauthorized.New("Session not found", errors.Options{
			HumanText: "Сессия не найдена",
		})
	}

	// Проверяем, не истекла ли сессия
	if session.ExpiresAt.Before(time.Now()) {
		return newTokens, errors.Unauthorized.New("Session ended", errors.Options{
			HumanText: "Истек строк действия токена авторизации, необходимо авторизоваться снова",
		})
	}

	// Удаляем все сессии для пользователя и устройства в бд
	if err = s.authRepository.DeleteSession(ctx, req.Necessary.UserID, req.Necessary.DeviceID); err != nil {
		return newTokens, err
	}

	// Создаем новую сессию и получаем новые токены
	newTokens.AccessToken, newTokens.RefreshToken, err = s.createSession(ctx, req.Necessary.UserID, req.Necessary.DeviceID)
	if err != nil {
		return newTokens, err
	}

	return newTokens, nil
}
