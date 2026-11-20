package users_service

import (
	"context"

	"github.com/ilez0660/Golangclicker/Server/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	GetUserByID(
		ctx context.Context,
		user_id int,
	) (domain.User, error)
	
	GetUserByEmail(
		ctx context.Context,
		email string,
	) (domain.User, error)
}

func NewUsersService(
	usersRepository UsersRepository,
) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (s *UsersService) GetUserByID(
	ctx context.Context,
	userID int,
) (domain.User, error) {
	return s.usersRepository.GetUserByID(ctx, userID)
}

func (s *UsersService) GetUserByEmail(
	ctx context.Context,
	email string,
) (domain.User, error) {
	return s.usersRepository.GetUserByEmail(ctx, email)
}
