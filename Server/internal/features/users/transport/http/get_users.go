package users_transport_http

import (
	"net/http"

	core_logger "github.com/ilez0660/Golangclicker/Server/internal/core/logger"
	core_http_response "github.com/ilez0660/Golangclicker/Server/internal/core/transport/http/response"
	core_http_utils "github.com/ilez0660/Golangclicker/Server/internal/core/transport/http/utils"
)

type GetUserResponse struct {
	ID       int     `json:"id"`
	Version  int     `json:"version"`
	FullName string  `json:"full_name"`
	Email    *string `json:"email"`
}

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
	
	userID, ok := core_http_utils.GetUserIDFromSession(r)
	if !ok {
		responseHandler.JSONResponse(
			map[string]string{"error": "не авторизован"},
			http.StatusUnauthorized,
		)
		return
	}

	userDomain, err := h.usersService.GetUserByID(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	response := GetUserResponse{
		ID:       userDomain.ID,
		Version:  userDomain.Version,
		FullName: userDomain.FullName,
		Email:    userDomain.Email,
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}