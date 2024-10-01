package handler

import (
	userAPI "github.com.damirqa.gophermart/internal/handler/api/user"
	"github.com.damirqa.gophermart/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux, useCases *usecase.UseCases) {
	router.Route("/api/user", func(r chi.Router) {
		r.Post("/register", userAPI.Register(useCases.UserRegisterUseCase))
		r.Post("/login", userAPI.Login())
		r.Post("/orders", userAPI.Orders())
		r.Get("/orders", userAPI.GetOrders())

		r.Get("/balance", userAPI.GetBalance())
		r.Post("/balance/withdraw", userAPI.Withdraw())
		r.Get("/withdrawals", userAPI.GetWithdrawals())
	})
}
