package interfaces

import (
	"user/service/pkg/domain"
	"user/service/pkg/models"
)

type UserUseCase interface {
	UsersSignUp(user models.UserSignUp) (domain.TokenUser, error)
	UsersLogin(user models.UserLogin) (domain.TokenUser, error)
}