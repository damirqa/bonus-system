package app

import (
	"context"
	"github.com.damirqa.gophermart/internal/config"
	"github.com.damirqa.gophermart/internal/domain/user/repository"
	"github.com.damirqa.gophermart/internal/domain/user/repository/pg"
	"github.com.damirqa.gophermart/internal/domain/user/service"
	"github.com.damirqa.gophermart/internal/infrastructure/logging"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

type Container struct {
	pool *pgxpool.Pool

	userRepository repository.UserRepositoryInterface
	userService    service.BaseUserServiceInterface
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Init() {
	c.pool = initPgxPool()

	initDomains(c)
}

func initDomains(c *Container) {
	c.userRepository, c.userService = initUserDomain(c)
}

func initPgxPool() *pgxpool.Pool {
	pgConfig, err := pgxpool.ParseConfig(config.GetDatabaseDSN())
	if err != nil {
		logging.GetLogger().Fatal("Failed to parse database DSN", zap.Error(err))
	}

	pgConfig.MaxConns = 20                       // = config.GetDatabaseMaxConnections()
	pgConfig.MinConns = 5                        // = config.GetDatabaseMinConnections()
	pgConfig.MaxConnLifetime = 3 * time.Minute   // = config.GetDatabaseMaxConnLifetime()
	pgConfig.MaxConnIdleTime = 1 * time.Minute   // = config.GetDatabaseMaxConnIdleTime()
	pgConfig.HealthCheckPeriod = 1 * time.Minute // = config.GetDatabaseHealthCheckPeriod()

	pgxPoolWithConfig, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		logging.GetLogger().Fatal("Failed to create PGX pool", zap.Error(err))
	}

	return pgxPoolWithConfig
}

func initUserDomain(c *Container) (repository.UserRepositoryInterface, *service.UserService) {
	userRepository := pg.NewUserPgRepository(c.pool)
	userService := service.NewUserService(userRepository)

	return userRepository, userService
}
