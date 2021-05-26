package infrastructure

import "github.com/robfig/cron"

type Cron struct {
	Cron *cron.Cron
}

func NewCron() *Cron {
	c := cron.New()
	return &Cron{
		Cron: c,
	}
}

func InitCron(cron *Cron) {
	cron.Cron.Start()
}
