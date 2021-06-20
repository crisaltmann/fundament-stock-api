package result_importer_service

import result_importer_domain "github.com/crisaltmann/fundament-stock-api/pkg/result_importer/domain"

func GetWegeTemplate() result_importer_domain.Template {
	dre := result_importer_domain.DRETemplate{
		Name:            "Dem. Result.",
		ReceitaLiquida: 5,
		ResultadoBruto: 7,
		DespesasVendas: 8,
		DespesasGeraisAdm: 9,
		LucroLiquido:   26,
	}

	balanco := result_importer_domain.BalancoTemplate{
		Name:                    "Bal. Patrim.",
		EmprestimoCirculante:    32,
		EmprestimoNaoCirculante: 39,
		Caixa:                   5,
	}

	return result_importer_domain.Template{
		Trimestre: 2,
		DRE: dre,
		Balanco: balanco,
	}
}
