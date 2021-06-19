package insight_service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/internal"
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	holding_domain "github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"
	insight_domain "github.com/crisaltmann/fundament-stock-api/pkg/insghts/domain"
	quarter_domain "github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
	"log"
	"strconv"
)

type Service struct {
	repository 	Repository
	quarterService QuarterService
	db *sql.DB
}

type Repository interface {
	DeleteByUser(ctx context.Context, tx *sql.Tx, idUser int64) error
	//GetQuarter(id int64) (quarter_domain.Trimestre, error)
	//GetQuarters() ([]quarter_domain.Trimestre, error)
}

type QuarterService interface {
	GetQuarter(id int64) (quarter_domain.Trimestre, error)
	GetQuarters() ([]quarter_domain.Trimestre, error)
}

func NewService(repository Repository, quarterService QuarterService, db *sql.DB) Service {
	return Service{
		repository: repository,
		quarterService: quarterService,
		db: db,
	}
}

func (s Service) CalculateInsights(ctx context.Context, holdings holding_domain.Holdings) error {
	/**
		montar map com chave: idAtivo-trimestre

		Terei todos os ativos por trimestre.

	 */
	var user int64
	insights := make([]insight_domain.Insight, 0)
	trimestreMap := make(map[string]holding_domain.HoldingAtivo)
	for i := 0; i < len(holdings.Holdings); i++ {
		holding := holdings.Holdings[i]
		for j := 0; j < len(holding.HoldingsAtivo); j++ {
			user = holding.HoldingsAtivo[j].Usuario
			holdingAtivo := holding.HoldingsAtivo[j]
			trimestreMap[buildKey(holdingAtivo.Ativo, holdingAtivo.Trimestre)] = holdingAtivo
		}
	}

	/**
		Busca o ativo do trimestre anterior e calcula a diferença.
	    Se não existir no trimestre anterior, salva 0%

	 */

	quarters, err := s.quarterService.GetQuarters()
	if err != nil {
		log.Print("Ocorreu um erro ao buscar os trimestres")
		return err
	}

	for _, holdingAtivo := range trimestreMap {
		currentQuarter := holdingAtivo.Trimestre
		lastQuarter := getLastQuarter(currentQuarter, quarters)
		lastQuarterHolding := trimestreMap[buildKey(holdingAtivo.Ativo, lastQuarter.Id)]

		var delta float32

		if lastQuarter.Id > 0 {
			delta = internal.RoundFloat(float32(holdingAtivo.ReceitaLiquida) - float32(lastQuarterHolding.ReceitaLiquida)/float32(lastQuarterHolding.ReceitaLiquida))
		} else {
			delta = 0
		}
		insights = append(insights, insight_domain.Insight{
			Id:           0,
			IdTrimestre:  currentQuarter,
			Usuario:    holdingAtivo.Usuario,
			IdAtivo:      holdingAtivo.Ativo.Id,
			ReceitaDelta: delta,
		})
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("Erro ao iniciar transacao")
	}
	//deleta todos insights
	err = s.repository.DeleteByUser(ctx, tx, user)
	if err != nil {
		log.Print("Erro ao deletar insights")
		tx.Rollback()
		return err
	}
	//insere novos insights


	return fmt.Errorf("tmp")
}

func getLastQuarter(quarter int64, quarters []quarter_domain.Trimestre) quarter_domain.Trimestre {
	return quarter_domain.Trimestre{}
}

func buildKey(ativo asset_domain.Asset, trimestre int64) string {
	return ativo.Codigo + strconv.FormatInt(trimestre, 10)
}