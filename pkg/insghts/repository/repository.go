package insight_repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/insghts/domain"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"time"
)

type Repository struct {
	DB *sql.DB
	cache 	 *cache.Cache
}

func NewRepository(db *sql.DB) Repository {
	cache := cache.New(1*time.Hour, 10*time.Minute)
	return Repository{DB: db, cache: cache}
}

func InitCache(r Repository) {
	//quarters, _ := r.GetQuarters()
	//for _, quarter := range quarters {
	//	r.cache.Add(strconv.FormatInt(quarter.Id, 10), quarter, cache.DefaultExpiration)
	//}
}

func (r Repository) DeleteByUser(ctx context.Context, tx *sql.Tx, idUser int64) error {
	prepare, err := tx.Prepare("DELETE FROM insights WHERE id_usuario = $1")

	if err != nil {
		log.Print("Ocorreu um erro ao preparar query de delete de insights por user")
		return err
	}

	defer prepare.Close()

	_, err = prepare.ExecContext(ctx, idUser)
	if err != nil {
		err = fmt.Errorf("Erro ao executar delete insights por user", err)
		return err
	}
	return nil
}

func (r Repository) SaveInsights(ctx context.Context, tx *sql.Tx, insight insight_domain.Insight) error {
	prepare, err := tx.Prepare("INSERT INTO insights (id_usuario, id_trimestre, id_ativo, " +
		" delta_receita_liquida, delta_ebitda, delta_lucro_liquido, delta_divida_liquida) VALUES($1, $2, $3, $4, $5, $6, $7);")

	if err != nil {
		log.Print("Ocorreu um erro ao preparar query de insert de insight")
		return err
	}

	defer prepare.Close()

	_, err = prepare.ExecContext(ctx, insight.Usuario, insight.IdTrimestre, insight.IdAtivo, insight.ReceitaDelta, insight.EbitdaDelta,
		insight.LucroDelta, insight.DividaDelta)
	if err != nil {
		err = fmt.Errorf("Erro ao executar insert de insight", err)
		return err
	}
	return nil
}