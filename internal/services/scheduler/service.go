package scheduler

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"

	"server/pkg/errors"
	"server/pkg/log"
	"server/internal/services"
	"server/internal/services/settings/model"
	settingsService "server/internal/services/settings/service"
)

const adminUser = 15

type Scheduler struct {
	cron            *cron.Cron
	settingsService *settingsService.Service
}

func NewScheduler(
	settingsService *settingsService.Service,

) *Scheduler {
	return &Scheduler{
		cron:            cron.New(),
		settingsService: settingsService,
	}
}

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
