package asset_sync

import (
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
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

func ConfigureJob(sync AssetSync, c *infrastructure.Cron) {
	cron := os.Getenv("STOCK_PRICE_UPDATE_CRON")
	if cron == "" {
		log.Print("Configurando cron com valor default.")
		c.Cron.AddFunc("0 0 2 * * *", sync.executeJob)
	} else {
		log.Print("Configurando cron com valor de variavel: " + cron)
		c.Cron.AddFunc(cron, sync.executeJob)
	}
}

func (a AssetSync) executeJob() {
	log.Println("Iniciando job de atualização de preços.")
	a.Service.UpdateAssetPrice()
	log.Println("Finalizando job de atualização de preços.")
}

