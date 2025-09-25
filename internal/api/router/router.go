package router

import (
	_ "github.com/LekcRg/steam-inventory/docs" // swaggo docs
	"github.com/LekcRg/steam-inventory/internal/api"
	"github.com/LekcRg/steam-inventory/internal/api/handlers"
	"github.com/LekcRg/steam-inventory/internal/api/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type Router struct {
	log zap.Logger
}

func New(h *handlers.Handlers, m *middlewares.Middlewares) *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		m.RequestLogger,
		middleware.CleanPath,
		middleware.AllowContentType("application/json"),
	)

	r.Get(api.PathAuthRedirect, h.AuthRedirect)
	r.Get(api.PathValidAuth, h.AuthValid)

	r.Group(func(r chi.Router) {
		r.Use(m.Auth)
		// r.Get(api.PathHi, h.Hi)
		r.Get(api.PathMe, h.UserInfo)
	})

	r.Get(api.PathSwagger, httpSwagger.Handler(
		httpSwagger.URL(api.PathSwaggerJSON),
	))

	return r
}
