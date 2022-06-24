package auth

import "summershare/internal/entities/web"

type AuthService interface {
	Login(username string, password string) web.Response
	Register(username string, email string, password string) web.Response
	RefreshToken(id string, refreshToken string) web.Response
}
