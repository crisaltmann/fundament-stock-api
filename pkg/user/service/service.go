package user_service

import (
	"fmt"
	user_domain "github.com/crisaltmann/fundament-stock-api/pkg/user/domain"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
)

type Service struct {
	repository         Repository
}

type Repository interface {
	Login(email string, password string) (int64, error)
}

func NewService(repository Repository) Service {
	return Service{
		repository:                     repository,
	}
}

func (s Service) Login(email string, password string) (user_domain.Login, error) {
	login, err := s.repository.Login(email, password)
	if err != nil {
		log.Print("Erro ao executar login")
		return user_domain.Login{}, err
	}

	if login <= 0 {
		return user_domain.Login{}, fmt.Errorf("User nao encontrado.")
	}
	myuuid := uuid.NewV4()

	return user_domain.Login{
		IdUser: login,
		Token:  myuuid.String(),
	}, nil
}