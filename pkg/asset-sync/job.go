package asset_sync

import (
	"github.com/robfig/cron"
	"log"
	"os"
)

type AssetSync struct {
	Service JobService
}

func NewAssetSync(service JobService) AssetSync {
	return AssetSync{
		Service: service,
	}
}

func ConfigureJob(sync AssetSync) {
	c := cron.New()

	cron := os.Getenv("STOCK_PRICE_UPDATE_CRON")
	if cron == "" {
		c.AddFunc("0 0 2 * * *", sync.executeJob)
	} else {
		c.AddFunc(cron, sync.executeJob)
	}

	c.Start()
}

func (a AssetSync) executeJob() {
	log.Println("Iniciando job de atualização de preços.")
	a.Service.updateAssetPrice()
	log.Println("Finalizando job de atualização de preços.")
}

