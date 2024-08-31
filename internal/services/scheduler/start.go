package scheduler

import (
	"context"
	"time"

	"server/internal/services"
	"server/internal/services/settings/model"
	"server/pkg/errors"
	"server/pkg/log"
)

func (s *Scheduler) Start() error {

	// Обновление валют
	_, err := s.cron.AddFunc("@daily", func() { // Every day at 00:00 UTC
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		if err := s.settingsService.UpdateCurrencies(ctx, model.UpdateCurrenciesReq{
			Necessary: services.NecessaryUserInformation{
				UserID:   adminUser,
				DeviceID: "system",
			},
		}); err != nil {
			log.Error(context.Background(), err)
		}
	})
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	s.cron.Start()

	return nil
}
