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

	quarterlyItem := asset_domain.AssetQuarterlyResult{
		Id:             0,
		Trimestre:      1,
		Ativo:          0,
		ReceitaLiquida: 3714000000,
		Ebitda:         0,
		MargemEbitda:   0,
		LucroLiquido:   0,
		MargemLiquida:  0,
		DividaLiquida:  0,
		DivEbitda:      0,
	}

	receitaPercentual := holding_service.CalcularFundamentos(item, quarterlyItem)
	expected := 138
	assert.EqualValues(t, expected, receitaPercentual)
}
