package insight_repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/insights/domain"
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

func (r Repository) GetInsights(usuario int64) ([]insight_domain.Insight, error) {
	rows, err := r.DB.Query("SELECT id, id_trimestre, id_usuario, id_ativo, delta_receita_liquida, delta_ebitda, delta_lucro_liquido, delta_divida_liquida " +
		" FROM insights WHERE id_usuario = $1 ", usuario)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de insights", err)
		return nil, err
	}
	defer rows.Close()
	insights := []insight_domain.Insight{}
	for rows.Next() {
		item := insight_domain.Insight{}
		err := rows.Scan(&item.Id, &item.IdTrimestre, &item.Usuario, &item.IdAtivo, &item.ReceitaDelta, &item.EbitdaDelta,
			&item.LucroDelta, &item.DividaDelta)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca de insights", err)
			return nil, err
		}
		insights = append(insights, item)
	}
	return insights, nil
}