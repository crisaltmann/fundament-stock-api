package asset_sync

import (
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/robfig/cron"
	"log"
)

type AssetSync struct {
	Service *Service
}

func configureJob(config *config.Config, sync *AssetSync) {
	c := cron.New()
	//c.AddFunc(config.Job.AssetSync.Cron, sync.executeJob)
	c.AddFunc("0 0/2 * * * *", sync.executeJob)

	c.Start()
}

func (a AssetSync) executeJob() {
	log.Println("Iniciando job de atualização de preços.")
	a.Service.updateAssetPrice()
	log.Println("Finalizando job de atualização de preços.")
}

