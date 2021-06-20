package quarter_service

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
	"github.com/rs/zerolog/log"
	"time"
)

type Service struct {
	repository 	Repository
}

type Repository interface {
	GetQuarter(id int64) (quarter_domain.Trimestre, error)
	GetQuarters() ([]quarter_domain.Trimestre, error)
	SaveQuarter(trimestre quarter_domain.Trimestre) (bool, error)
	GetQuarterByCode(code string) (quarter_domain.Trimestre, error)
}

func NewService(repository Repository) Service {
	return Service{
		repository: repository,
	}
}

func (s Service) GetQuarter(id int64) (quarter_domain.Trimestre, error) {
	return s.repository.GetQuarter(id)
}

func (s Service) GetQuarters() ([]quarter_domain.Trimestre, error) {
	trimestres, err := s.repository.GetQuarters()
	if err != nil {
		return nil, err
	}
	return trimestres, nil
}

func (s Service) GetQuarterByDate(date time.Time) (quarter_domain.Trimestre, error) {
	trimestres, err := s.GetQuarters()
	if err != nil {
		log.Print("Erro ao buscar trimestres")
		return quarter_domain.Trimestre{}, err
	}

	for _, trimestre := range trimestres {
		if (date.Before(trimestre.DataFim) || date.Equal(trimestre.DataFim)) && (date.After(trimestre.DataInicio) || date.Equal(trimestre.DataInicio)){
			return trimestre, nil
		}
	}
	return quarter_domain.Trimestre{}, fmt.Errorf("Trimestre nao encontrado")
}

func (s Service) CreateQuarterByDate(date time.Time) (bool, error) {
	ano, mes, _ := date.Date()
	trimestre := getTrimestre(int(mes))
	dataInicio, dataFim, err := getDataInicioFim(trimestre, ano)
	if err != nil {
		return false, err
	}
	lastQuarter, err := s.getLastQuarter(ano, trimestre)
	if err != nil {
		return false, err
	}
	savedQuarter := quarter_domain.Trimestre{
		Id:                0,
		Codigo:            getCodigo(ano, trimestre),
		Ano:               ano,
		Trimestre:         trimestre,
		DataInicio:        dataInicio,
		DataFim:           dataFim,
		TrimestreAnterior: lastQuarter.Id,
	}

	save, err := s.repository.SaveQuarter(savedQuarter)
	if save && err == nil {
		//force adding to cache
		s.repository.GetQuarterByCode(savedQuarter.Codigo)
	}
	return save, err
}

func (s Service) getLastQuarter(ano int, trimestre int) (quarter_domain.Trimestre, error) {
	usedTrimestre := trimestre
	usedAno := ano
	if trimestre == 1 {
		usedTrimestre = 4
		usedAno--
	} else {
		usedTrimestre--
	}

	quarters, err  := s.GetQuarters()
	if err != nil {
		return quarter_domain.Trimestre{}, err
	}
	for _, quarter := range quarters {
		if quarter.Trimestre == usedTrimestre && quarter.Ano == usedAno {
			return quarter, nil
		}
	}
	return quarter_domain.Trimestre{}, nil
}

func getCodigo(ano int, trimestre int) string {
	return fmt.Sprintf("%04d_%02d", ano, trimestre)
}

func getTrimestre(mes int) int {
	switch mes {
	case 1,2,3:
		return 1
	case 4,5,6:
		return 2
	case 7,8,9:
		return 3
	case 10, 11, 12:
		return 4
	default:
		return 0
	}
}

func getDataInicioFim(trimestre int, ano int) (time.Time, time.Time, error) {
	layout := "02/01/2006"

	switch trimestre {
	case 1:
		dateInit, err := time.Parse(layout, "01/01/" + fmt.Sprintf("%d", ano))
		if err != nil {
			log.Print("Erro ao fazer o parser da data %s")
			return time.Time{}, time.Time{}, err
		}
		dateFim, err := time.Parse(layout, "31/03/" + fmt.Sprintf("%d", ano))
		if err != nil {
			log.Print("Erro ao fazer o parser da data %s")
			return time.Time{}, time.Time{}, err
		}
		return dateInit, dateFim, nil
	case 2:
		dateInit, err := time.Parse(layout, "01/04/" + fmt.Sprintf("%d", ano))
		if err != nil {
			log.Print("Erro ao fazer o parser da data %s")
			return time.Time{}, time.Time{}, err
		}
		dateFim, err := time.Parse(layout, "30/06/" + fmt.Sprintf("%d", ano))
		if err != nil {
			log.Print("Erro ao fazer o parser da data %s")
			return time.Time{}, time.Time{}, err
		}
		return dateInit, dateFim, nil
	case 3:
		dateInit, err := time.Parse(layout, "01/07/" + fmt.Sprintf("%d", ano))
		if err != nil {
			log.Print("Erro ao fazer o parser da data %s")
			return time.Time{}, time.Time{}, err
		}
		dateFim, err := time.Parse(layout, "30/09/" + fmt.Sprintf("%d", ano))
		if err != nil {
			log.Print("Erro ao fazer o parser da data %s")
			return time.Time{}, time.Time{}, err
		}
		return dateInit, dateFim, nil
	case 4:
		dateInit, err := time.Parse(layout, "01/10/" + fmt.Sprintf("%d", ano))
		if err != nil {
			log.Print("Erro ao fazer o parser da data %s")
			return time.Time{}, time.Time{}, err
		}
		dateFim, err := time.Parse(layout, "31/12/" + fmt.Sprintf("%d", ano))
		if err != nil {
			log.Print("Erro ao fazer o parser da data %s")
			return time.Time{}, time.Time{}, err
		}
		return dateInit, dateFim, nil
	default:
		return time.Time{}, time.Time{}, fmt.Errorf("Não foi possivel converter a data.")
	}
}