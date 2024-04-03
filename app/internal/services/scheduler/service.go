package scheduler

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"

	adminService "server/app/internal/services/admin/service"
	"server/pkg/errors"
	"server/pkg/logging"
)

type Scheduler struct {
	cron         *cron.Cron
	adminService *adminService.Service
	logger       *logging.Logger
}

func NewScheduler(
	adminService *adminService.Service,
	logger *logging.Logger,
) *Scheduler {
	return &Scheduler{
		cron:         cron.New(),
		adminService: adminService,
		logger:       logger,
	}
}

func (s *Scheduler) Start() error {

	// Обновление валют
	_, err := s.cron.AddFunc("@daily", func() { // Every day at 00:00 UTC
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		if err := s.adminService.UpdateCurrencies(ctx); err != nil {
			s.logger.Error(err)
		}
	})
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	s.cron.Start()

	return nil
}
