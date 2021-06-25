package user_api

type LoginRequest struct {
	Email 		string		`json:"email"`
	Password	string 		`json:"password"`
}

type LoginResponse struct {
	IdUsuario	int64		`json:"id_usuario"`
	Token		string		`json:"token"`
}