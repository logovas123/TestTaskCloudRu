package db

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"rl-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

// создаём новый пул соединений к базе, данные для подключения к базе берутся из конфига
func NewConnPostgres(ctx context.Context, cfg *config.DBConfig) (*pgxpool.Pool, error) {
	dbcfg, err := pgxpool.ParseConfig(CreateDataSourceName(cfg))
	if err != nil {
		return nil, fmt.Errorf("error parsing of config: %w", err)
	}

	dbcfg.MinConns = 1

	dbcfg.ConnConfig.ConnectTimeout = 15 * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, dbcfg)
	if err != nil {
		return nil, fmt.Errorf("error create connect to db: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("error ping to db: %w", err)
	}

	slog.Info("connect to db create success")

	return pool, nil
}

func CreateDataSourceName(cfg *config.DBConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
}
