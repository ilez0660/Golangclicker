package users_transport_http

import (
	"net/http"

	"github.com/ilez0660/Golangclicker/Server/internal/core/domain"
	core_logger "github.com/ilez0660/Golangclicker/Server/internal/core/logger"
	core_http_request "github.com/ilez0660/Golangclicker/Server/internal/core/transport/http/request"
	core_http_response "github.com/ilez0660/Golangclicker/Server/internal/core/transport/http/response"
	core_http_utils "github.com/ilez0660/Golangclicker/Server/internal/core/transport/http/utils"
)

type CreateUserRequest struct {
	FullName string `json:"full_name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
}

type CreateUserResponse struct {
	ID       int     `json:"id"`
	Version  int     `json:"version"`
	FullName string  `json:"full_name"`
	Email    *string `json:"email"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}
	userDomain := domainFromDTO(request)

	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")

		return
	}

	core_http_utils.CreateSession(rw, userDomain.ID)

	response := dtoFromDomain(userDomain)
	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, &dto.Email)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:       user.ID,
		Version:  user.Version,
		FullName: user.FullName,
		Email:    user.Email,
	}

}
