package insight_service

type Service struct {
	repository 	Repository
}

type Repository interface {
	//GetQuarter(id int64) (quarter_domain.Trimestre, error)
	//GetQuarters() ([]quarter_domain.Trimestre, error)
}

func NewService(repository Repository) Service {
	return Service{
		repository: repository,
	}
}

//func (s Service) GetQuarter(id int64) (quarter_domain.Trimestre, error) {
//	return s.repository.GetQuarter(id)
//}
//
//func (s Service) GetQuarters() ([]quarter_domain.Trimestre, error) {
//	trimestres, err := s.repository.GetQuarters()
//	if err != nil {
//		return nil, err
//	}
//	return trimestres, nil
//}