package holding_service_test

import (
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	holding_service "github.com/crisaltmann/fundament-stock-api/pkg/holding/service"
	portfolio_domain "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetExpeditionContainer_Success(t *testing.T) {
	ativo := portfolio_domain.Ativo{
		Id:      0,
		Codigo:  "WEGE3",
		Logo:    "",
		Total:   4197317998,
		Cotacao: 0,
	}
	item := portfolio_domain.Portfolio{
		Ativo:      ativo,
		Quantidade: 157,
		Valor:      0,
		Usuario:    0,
	}
	// quantidade / total = 0,000000037404838

	quarterlyItem := asset_domain.AssetQuarterlyResult {
		Id:             0,
		Trimestre:      1,
		Ativo:          0,
		ReceitaLiquida: 3714000000,
		Ebitda:         619000000,
		MargemEbitda:   0,
		LucroLiquido:   454000000,
		MargemLiquida:  0,
		DividaLiquida:  454000000,
		DivEbitda:      0,
	}

	receitaLiquida, ebitda, lucroLiquido, divida := holding_service.CalcularFundamentos(item, quarterlyItem)
	assert.EqualValues(t, 138, receitaLiquida)
	assert.EqualValues(t, 23, ebitda)
	assert.EqualValues(t, 16, lucroLiquido)
	assert.EqualValues(t, 16, divida)
}
