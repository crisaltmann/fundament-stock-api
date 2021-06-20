package result_importer_service

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	result_importer_domain "github.com/crisaltmann/fundament-stock-api/pkg/result_importer/domain"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

func ImportWege() (result_importer_domain.ImporterResults, error) {
	f, err := excelize.OpenFile("balanco-WEGE.xlsx")
	if err != nil {
		log.Print("Erro ao abrir arquivo")
		return result_importer_domain.ImporterResults{}, err
	}

	template := GetWegeTemplate()

	rowsBalanco, err := f.GetRows(template.Balanco.Name)
	if err != nil {
		log.Print("Erro ao ler rows de balanco")
		return result_importer_domain.ImporterResults{}, err
	}

	rowsDre, err := f.GetRows(template.DRE.Name)
	if err != nil {
		log.Print("Erro ao ler rows de dre")
		return result_importer_domain.ImporterResults{}, err
	}

	trimestres, err := getTrimestres(rowsBalanco)
	if err != nil {
		log.Print("Erro ao ler trimestres")
		return result_importer_domain.ImporterResults{}, err
	}

	results := result_importer_domain.ImporterResults{}
	results.Codigo = "WEGE3"
	for _, trimestre := range trimestres {
		balanco, err := getBalanco(template, trimestre, rowsBalanco)
		if err != nil {
			log.Print("Ocorreu um erro ao fazer o parser do balanco")
			return result_importer_domain.ImporterResults{}, err
		}
		dre, err := getDre(template, trimestre, rowsDre)
		if err != nil {
			log.Print("Ocorreu um erro ao fazer o parser da dre")
			return result_importer_domain.ImporterResults{}, err
		}
		result := result_importer_domain.ImporterResult{
			Trimestre: trimestre,
			Balanco:   balanco,
			DRE:       dre,
		}
		results.Results = append(results.Results, result)
	}
	return results, nil
}

func getTrimestres(rows [][]string) ([]result_importer_domain.Trimestre, error) {
	trimestres := make([]result_importer_domain.Trimestre, 0)

	trimestresRow := rows[1]
	for i := 1; i < len(trimestresRow); i++ {
		value := trimestresRow[i]

		date, err := time.Parse("02/01/2006", value)
		if err != nil {
			log.Print("Erro ao fazer o parser da data %s")
			return nil, err
		}

		trimestre := result_importer_domain.Trimestre{
			Y:     i,
			Value: value,
			Date: date,
		}
		trimestres = append(trimestres, trimestre)
	}
	return trimestres, nil
}

func getBalanco(template result_importer_domain.Template, trimestre result_importer_domain.Trimestre, rows [][]string) (result_importer_domain.Balanco, error) {
	emprestimoCirculante, err := strconv.ParseFloat(rows[template.Balanco.EmprestimoCirculante-1][trimestre.Y], 32)
	if err != nil {
		return result_importer_domain.Balanco{},err
	}
	emprestimoNaoCirculante, err := strconv.ParseFloat(rows[template.Balanco.EmprestimoNaoCirculante-1][trimestre.Y], 32)
	if err != nil {
		return result_importer_domain.Balanco{},err
	}
	caixa, err := strconv.ParseFloat(rows[template.Balanco.Caixa-1][trimestre.Y], 32)
	if err != nil {
		return result_importer_domain.Balanco{},err
	}

	return result_importer_domain.Balanco{
		EmprestimoCirculante:    float32(emprestimoCirculante),
		EmprestimoNaoCirculante: float32(emprestimoNaoCirculante),
		Caixa:                   float32(caixa),
	}, nil
}

func getDre(template result_importer_domain.Template, trimestre result_importer_domain.Trimestre, rows [][]string) (result_importer_domain.DRE, error) {
	receitaliquida, err := strconv.ParseFloat(rows[template.DRE.ReceitaLiquida-1][trimestre.Y], 32)
	if err != nil {
		return result_importer_domain.DRE{},err
	}
	resultadobruto, err := strconv.ParseFloat(rows[template.DRE.ResultadoBruto-1][trimestre.Y], 32)
	if err != nil {
		return result_importer_domain.DRE{},err
	}
	despesasvendas, err := strconv.ParseFloat(rows[template.DRE.DespesasVendas-1][trimestre.Y], 32)
	if err != nil {
		return result_importer_domain.DRE{},err
	}
	despesasgerais, err := strconv.ParseFloat(rows[template.DRE.DespesasGeraisAdm-1][trimestre.Y], 32)
	if err != nil {
		return result_importer_domain.DRE{},err
	}
	lucroliquido, err := strconv.ParseFloat(rows[template.DRE.LucroLiquido-1][trimestre.Y], 32)
	if err != nil {
		return result_importer_domain.DRE{},err
	}
	return result_importer_domain.DRE{
		ReceitaLiquida:    float32(receitaliquida),
		ResultadoBruto:    float32(resultadobruto),
		DespesasVendas:    float32(despesasvendas),
		DespesasGeraisAdm: float32(despesasgerais),
		LucroLiquido:      float32(lucroliquido),
	}, nil
}
