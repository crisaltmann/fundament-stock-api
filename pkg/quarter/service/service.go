package quarter_service

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
)

type Service struct {
	Repository 	Repository
	Cache 		*cache.Cache
}

type Repository interface {
	GetQuarter(id int64) (quarter_domain.Trimestre, error)
	GetQuarters() ([]quarter_domain.Trimestre, error)
}

func NewService(repository Repository) Service {
	cache := cache.New(1*time.Minute, 1*time.Minute)
	return Service{
		Repository: repository,
		Cache: cache,
	}
}

func (s Service) GetQuarter(id int64) (quarter_domain.Trimestre, error) {
	trimestre, found :=  s.Cache.Get(strconv.FormatInt(id, 10))
	if found {
		return trimestre.(quarter_domain.Trimestre), nil
	}
	trimestreDB, err := s.Repository.GetQuarter(id)
	if err != nil {
		return quarter_domain.Trimestre{}, err
	}
	s.Cache.Add(strconv.FormatInt(id, 10), trimestreDB, cache.DefaultExpiration)
	return trimestreDB, nil
}

func (s Service) GetQuarters() ([]quarter_domain.Trimestre, error) {
	trimestres, err := s.Repository.GetQuarters()
	if err != nil {
		return nil, err
	}
	return trimestres, nil
}