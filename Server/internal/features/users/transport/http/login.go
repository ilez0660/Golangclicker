package users_transport_http

import (
	"net/http"

	core_logger "github.com/ilez0660/Golangclicker/Server/internal/core/logger"
	core_http_request "github.com/ilez0660/Golangclicker/Server/internal/core/transport/http/request"
	core_http_response "github.com/ilez0660/Golangclicker/Server/internal/core/transport/http/response"
	core_http_utils "github.com/ilez0660/Golangclicker/Server/internal/core/transport/http/utils"
)

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (h *UsersHTTPHandler) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request LoginRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain, err := h.usersService.GetUserByEmail(ctx, request.Email)
	if err != nil {
		responseHandler.JSONResponse(map[string]string{"error": "пользователь не найден"}, http.StatusUnauthorized)
		return
	}

	core_http_utils.CreateSession(rw, userDomain.ID)

	response := dtoFromDomain(userDomain)
	responseHandler.JSONResponse(response, http.StatusOK)
}
