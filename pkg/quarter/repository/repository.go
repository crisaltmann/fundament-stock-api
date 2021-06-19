package quarter_repository

import (
	"database/sql"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
	"github.com/patrickmn/go-cache"
	"strconv"
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
	quarters, _ := r.GetQuarters()
	for _, quarter := range quarters {
		r.cache.Add(strconv.FormatInt(quarter.Id, 10), quarter, cache.DefaultExpiration)
	}
}

func (r Repository) GetQuarter(id int64) (quarter_domain.Trimestre, error) {
	trimestreCache, found :=  r.cache.Get(strconv.FormatInt(id, 10))
	if found {
		return trimestreCache.(quarter_domain.Trimestre), nil
	}

	rows, err := r.DB.Query("SELECT id, codigo, ano, trimestre, data_inicio, data_fim FROM trimestre WHERE id = $1", id)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de trimestre", err)
		return quarter_domain.Trimestre{}, err
	}
	defer rows.Close()
	trimestre := quarter_domain.Trimestre{}
	for rows.Next() {
		err := rows.Scan(&trimestre.Id, &trimestre.Codigo, &trimestre.Ano, &trimestre.Trimestre, &trimestre.DataInicio, &trimestre.DataFim)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca do trimestre", err)
			return quarter_domain.Trimestre{}, err
		}
	}

	if err != nil {
		r.cache.Add(strconv.FormatInt(id, 10), trimestre, cache.DefaultExpiration)
	}

	return trimestre, nil
}

func (r Repository) GetQuarters() ([]quarter_domain.Trimestre, error) {
	if r.cache.ItemCount() > 0 {
		trimestres := make([]quarter_domain.Trimestre, 0)
		for _, trim := range r.cache.Items() {
			trimestres = append(trimestres, trim.Object.(quarter_domain.Trimestre))
		}
		return trimestres, nil
	}

	rows, err := r.DB.Query("SELECT id, codigo, ano, trimestre, data_inicio, data_fim FROM trimestre")
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de trimestres", err)
		return nil, err
	}
	defer rows.Close()
	trimestres := []quarter_domain.Trimestre{}
	for rows.Next() {
		trimestre := quarter_domain.Trimestre{}
		err := rows.Scan(&trimestre.Id, &trimestre.Codigo, &trimestre.Ano, &trimestre.Trimestre, &trimestre.DataInicio, &trimestre.DataFim)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca do trimestres", err)
			return nil, err
		}
		trimestres = append(trimestres, trimestre)
	}
	return trimestres, nil
}