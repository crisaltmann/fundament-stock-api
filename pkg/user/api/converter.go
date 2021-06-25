package user_api

import (
	user_domain "github.com/crisaltmann/fundament-stock-api/pkg/user/domain"
)

func convertDomainToDto(login user_domain.Login) LoginResponse {
	return LoginResponse{
		IdUsuario: login.IdUser,
		Token:     login.Token,
	}
}