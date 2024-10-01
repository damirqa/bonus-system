package user

import (
	"encoding/json"
	"errors"
	CustomErrors "github.com.damirqa.gophermart/internal/errs"
	"github.com.damirqa.gophermart/internal/infrastructure/logging"
	"github.com.damirqa.gophermart/internal/usecase/user/auth"
	"go.uber.org/zap"
	"net/http"
)

type RegisterCase struct {
}

type ApiUserRegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func Register(u *auth.UserRegisterUseCase) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var registerRequest ApiUserRegisterRequest
		if err := json.NewDecoder(request.Body).Decode(&ApiUserRegisterRequest{}); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		err := u.Register(registerRequest.Login, registerRequest.Password)
		if err != nil {
			var userExistsErr *CustomErrors.ErrUserAlreadyExists
			if errors.As(err, &userExistsErr) {
				http.Error(writer, "Login or password already exist", http.StatusConflict)
				return
			}

			logging.GetLogger().Error("Error while registering user", zap.Error(err))
			http.Error(writer, "Error while registering user", http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}
