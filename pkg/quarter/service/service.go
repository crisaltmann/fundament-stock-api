package quarter_service

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
)

type Service struct {
	Repository Repository
}

type Repository interface {
	GetQuarter(id int64) (quarter_domain.Trimestre, error)
	GetQuarters() ([]quarter_domain.Trimestre, error)
}

func NewService(repository Repository) Service {
	return Service{Repository: repository}
}

func (s Service) GetQuarter(id int64) (quarter_domain.Trimestre, error) {
	trimestre, err := s.Repository.GetQuarter(id)
	if err != nil {
		return quarter_domain.Trimestre{}, err
	}
	return trimestre, nil
}

func (s Service) GetQuarters() ([]quarter_domain.Trimestre, error) {
	trimestres, err := s.Repository.GetQuarters()
	if err != nil {
		return nil, err
	}
	return trimestres, nil
}