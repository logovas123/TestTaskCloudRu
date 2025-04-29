package server

import (
	"context"
	"errors"
	"log/slog"

	"rl-service/internal/usecase/bucketUsecase"

	"github.com/jackc/pgx/v5"
)

type Client struct {
	IP    string
	Count uint64
	Rate  uint64
}

// функция создавалась в последний момент,
// нужна для восстановления бакетов при перезапуске сервиса
func (s *Server) RecoverBuckets(u *bucketUsecase.BucketUC) error {
	query := `
        SELECT ip::TEXT, count, rate FROM clients
    `

	ctx := context.Background()

	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var client Client
		if err := rows.Scan(&client.IP, &client.Count, &client.Rate); err != nil {
			return err
		}

		if err := u.StartBucket(ctx, toStartBucketRequest(client)); err != nil {
			slog.Error("error start bucket", "ip", client.IP, "error", err)
			continue
		}
	}

	slog.Info("recover buckets success")

	return nil
}

func toStartBucketRequest(
	req Client,
) bucketUsecase.StartBucketRequest {
	return bucketUsecase.StartBucketRequest{
		IP:        req.IP,
		Rate:      req.Rate,
		MaxTokens: req.Count,
	}
}
