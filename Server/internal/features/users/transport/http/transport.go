package users_transport_http

import (
	"context"
	"net/http"

	"github.com/ilez0660/Golangclicker/Server/internal/core/domain"
	core_http_server "github.com/ilez0660/Golangclicker/Server/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	GetUserByID(
		ctx context.Context,
		id int,
	) (domain.User, error)

	GetUserByEmail(
		ctx context.Context,
		email string,
	) (domain.User, error)
}

func NewUsersHTTPHandler(
	usersService UsersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/me", // Фронт стучится именно сюда
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users/login",
			Handler: h.Login,
		},
	}
}
