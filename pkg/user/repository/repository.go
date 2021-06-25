package user_repository

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	DB       *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) Login(email string, password string) (int64, error) {
	rows, err := r.DB.Query("SELECT id FROM usuario WHERE email = $1 and password = $2", email, password)
	defer rows.Close()

	if err != nil {
		err = fmt.Errorf("Erro ao executar busca de usuario", err)
		return 0, err
	}
	defer rows.Close()
	var idUser int64
	for rows.Next() {
		err := rows.Scan(&idUser)
		if err != nil {
			err = fmt.Errorf("Erro ao executar busca de usuario", err)
			return 0, err
		}
	}

	return idUser, nil
}