package result_importer_service

import (
	"errors"
	"fmt"
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
	result_importer_domain "github.com/crisaltmann/fundament-stock-api/pkg/result_importer/domain"
	"github.com/rs/zerolog/log"
	"sort"
	"time"
)

type Importer struct {
	quarterService 		QuarterService
	quarterlyResult 	QuarterlyResultService
	assetService		AssetService
}

type QuarterService interface {
	GetQuarterByDate(date time.Time) (quarter_domain.Trimestre, error)
	CreateQuarterByDate(date time.Time) (bool, error)
}

type AssetService interface {
	GetByCode(code string) (asset_domain.Asset, error)
	GetAllAssets() ([]asset_domain.Asset, error)
}

type QuarterlyResultService interface {
	ExistAssetQuarterlyResult(idAtivo int64, idTrimestre int64) (bool, error)
	InsertAssetQuarterlyResult(aqResult asset_domain.AssetQuarterlyResult) (bool, error)
}

func NewImporter(service QuarterService, quarterlyResult QuarterlyResultService, assetService AssetService) Importer {
	return Importer{
		quarterService: service,
		quarterlyResult: quarterlyResult,
		assetService: assetService,
	}
}

func (i Importer) StartImporterAll() error {
	assets, err := i.assetService.GetAllAssets()
	if err != nil {
		log.Printf("Erro ao buscar ativos")
		return err
	}
	for _, asset := range assets {
		log.Printf("Importando ativo %s", asset.Codigo)
		err = i.StartImporter(asset.Codigo)
		if err != nil {
			log.Printf("ERRO AO IMPORTAR ASSET.")
		}
	}
	return nil
}

func (i Importer) StartImporter(code string) error {

	var results result_importer_domain.ImporterResults
	var err error

	switch code {
	case "WEGE3", "HYPE3", "FRAS3", "MGLU3", "ABEV3":
		results, err = Import(code)
		if err != nil {
			log.Printf("Erro ao importar resultados de %s", code)
			return err
		}
	default:
		return fmt.Errorf("Ativo não suportado")
	}

	err = i.saveQuarters(results)
	if err != nil {
		log.Print("Erro ao salvar quarters")
		return err
	}

	asset, err := i.assetService.GetByCode(results.Codigo)
	if err != nil {
		log.Print("Erro ao buscar o ativo %s", results.Codigo)
		return err
	}

	for _, result := range results.Results {
		quarter, err := i.quarterService.GetQuarterByDate(result.Trimestre.Date)
		if err != nil {
			log.Print("Erro ao buscar o quarter")
			return err
		}

		existsQuarterly, err := i.existQuarterlyResult(asset.Id, quarter.Id)
		if err != nil {
			log.Print("Ocorreu um erro ao buscar o resultado de trimestre %d - %d", asset.Id, quarter.Id)
			return err
		}

		if !existsQuarterly {
			i.inserirResultado(asset, quarter, result)
		}
	}

	return nil
}

func (i Importer) saveQuarters(results result_importer_domain.ImporterResults) error {
	quarters := make([]time.Time, 0)
	for _, result := range results.Results {
		quarters = append(quarters, result.Trimestre.Date)
	}
	sort.Slice(quarters, func(i, j int) bool {
		return quarters[i].Before(quarters[j])
	})

	for _, quarter := range quarters {
		trim, err := i.quarterService.GetQuarterByDate(quarter)
		if err != nil && err.Error() != "Trimestre nao encontrado" {
			return err
		}

		if trim.Id > 0 {
			continue
		}

		_, err = i.quarterService.CreateQuarterByDate(quarter)
		if err != nil {
			log.Print("Erro ao salvar quarter.")
			return err
		}
	}
	return nil
}

func (i Importer) existQuarterlyResult(idAtivo int64, idTrimestre int64) (bool, error) {
	result, err := i.quarterlyResult.ExistAssetQuarterlyResult(idAtivo, idTrimestre)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (i Importer) inserirResultado(asset asset_domain.Asset, quarter quarter_domain.Trimestre, result result_importer_domain.ImporterResult) error {
	ebitda := result.DRE.ResultadoBruto - result.DRE.DespesasVendas - result.DRE.DespesasVendas
	divLiquida := result.Balanco.EmprestimoCirculante + result.Balanco.EmprestimoNaoCirculante - result.Balanco.Caixa

	aqResult := asset_domain.AssetQuarterlyResult{
		Trimestre:      quarter.Id,
		Ativo:          asset.Id,
		ReceitaLiquida: int64(result.DRE.ReceitaLiquida * 1000),
		Ebitda:         int64(ebitda * 1000),
		LucroLiquido:   int64(result.DRE.LucroLiquido * 1000),
		DividaLiquida:  int64(divLiquida * 1000),
	}

	salvo, err := i.quarterlyResult.InsertAssetQuarterlyResult(aqResult)
	if err != nil || !salvo {
		log.Print("Ocorreu um erro ao salvar quarterly result")
		return errors.New("Ocorreu um erro ao salvar quarterly result")
	}

	return nil
}