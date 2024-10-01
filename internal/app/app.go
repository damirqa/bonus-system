package app

import (
	"context"
	"github.com.damirqa.gophermart/internal/config"
	"github.com.damirqa.gophermart/internal/handler"
	"github.com.damirqa.gophermart/internal/infrastructure/logging"
	"github.com.damirqa.gophermart/internal/usecase"
	"github.com.damirqa.gophermart/internal/usecase/user/auth"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	container  *Container
	router     *chi.Mux
	httpServer *http.Server
	useCases   *usecase.UseCases
}

func Run() {
	app := NewApp()
	app.Init()
	app.Start()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	app.Shutdown()
}

func NewApp() *App {
	app := &App{}
	return app
}

func (app *App) Init() {
	logging.GetLogger().Info("Initializing app")

	app.initLogger()
	app.initContainer()
	app.initUseCases()
	app.initRouter()
	app.initHTTPServer()

	logging.GetLogger().Info("App initialized")
}

func (app *App) initLogger() {
	err := logging.Initialize(config.GetLogLevel())
	if err != nil {
		panic(err)
	}
}

func (app *App) initContainer() {
	app.container = NewContainer()
	app.container.Init()
}

func (app *App) initUseCases() {
	app.useCases = &usecase.UseCases{
		UserRegisterUseCase: auth.NewUserRegisterUseCase(app.container.userService),
	}
}

func (app *App) initRouter() {
	app.router = chi.NewRouter()
	handler.RegisterRoutes(app.router, app.useCases)
}

func (app *App) initHTTPServer() {
	app.httpServer = &http.Server{
		Addr:    config.GetAddress(),
		Handler: app.router,
	}
}

func (app *App) Start() {
	logging.GetLogger().Info("Starting server", zap.String("address", config.GetAddress()))

	err := http.ListenAndServe(config.GetAddress(), app.router)
	if err != nil {
		logging.GetLogger().Error("Failed to start server", zap.Error(err))
	}

	logging.GetLogger().Info("Server started")
}

func (app *App) Shutdown() {
	logging.GetLogger().Info("Shutting down server")

	err := app.httpServer.Shutdown(context.Background())
	if err != nil {
		logging.GetLogger().Error("Failed to stop server", zap.Error(err))
	}

	logging.GetLogger().Info("Server stopped")
}
