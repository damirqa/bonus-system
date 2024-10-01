package pg

import (
	"context"
	"errors"

	CustomErrors "github.com.damirqa.gophermart/internal/errs"
	"github.com.damirqa.gophermart/internal/infrastructure/logging"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

type UserPgRepository struct {
	pool *pgxpool.Pool
}

func NewUserPgRepository(pool *pgxpool.Pool) *UserPgRepository {
	return &UserPgRepository{pool: pool}
}

func (r *UserPgRepository) Insert(login, hashPassword string) error {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := r.pool.Acquire(ctxWithTimeout)
	if err != nil {
		return err
	}

	defer conn.Release()

	_, err = conn.Conn().Prepare(context.Background(), "users", "INSERT INTO users (login, hashPassword) VALUES ($1, $2)")
	if err != nil {
		logging.GetLogger().Error("Failed to prepare statement: %v", zap.Error(err))
		return err
	}

	_, err = conn.Conn().Exec(context.Background(), "users", login, hashPassword)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return CustomErrors.NewErrUserAlreadyExists(err)
		}

		logging.GetLogger().Error("Failed to prepare statement: %v", zap.Error(err))
		return err
	}

	return nil
}
