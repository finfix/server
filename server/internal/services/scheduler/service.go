package scheduler

import (
	"github.com/robfig/cron/v3"

	settingsService "server/internal/services/settings/service"
)

const adminUser = 15

type Scheduler struct {
	cron            *cron.Cron
	settingsService *settingsService.SettingsService
}

func NewScheduler(
	settingsService *settingsService.SettingsService,

) *Scheduler {
	return &Scheduler{
		cron:            cron.New(),
		settingsService: settingsService,
	}
}
