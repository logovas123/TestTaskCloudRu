package clientRepository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type CreateClientIfNotExistRequest struct {
	IP    string
	Count uint64
	Rate  uint64
}

type CreateClientIfNotExistResponse struct {
	IP      string
	Count   uint64
	Rate    uint64
	IsExist bool
}

func (r *ClientRepository) CreateClientIfNotExist(
	ctx context.Context,
	req CreateClientIfNotExistRequest,
) (CreateClientIfNotExistResponse, error) {
	query := `
        SELECT ip, count, rate FROM clients 
        WHERE ip = $1
    `

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		slog.Error("error creating tx", "error", err)
		return CreateClientIfNotExistResponse{}, err
	}
	defer func() {
		var e error
		if err == nil {
			e = tx.Commit(ctx)
		} else {
			e = tx.Rollback(ctx)
		}

		if err == nil && e != nil {
			err = fmt.Errorf("finishing transaction: %w", e)
		}
	}()

	var client CreateClientIfNotExistResponse
	err = tx.QueryRow(ctx, query, req.IP).
		Scan(&client.IP, &client.Count, &client.Rate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			client, err = insertClient(ctx, tx, req)
			if err != nil {
				slog.Error("error creating client in db", "error", err)
				return CreateClientIfNotExistResponse{}, err
			}
			return client, nil
		}
		slog.Error("error executing query to db", "error", err)
		return CreateClientIfNotExistResponse{}, err
	}

	client.IsExist = true

	return client, nil
}

func insertClient(
	ctx context.Context,
	tx pgx.Tx,
	req CreateClientIfNotExistRequest,
) (CreateClientIfNotExistResponse, error) {
	query := `
		INSERT INTO clients (ip, count, rate)
		VALUES ($1, $2, $3)
		RETURNING ip, count, rate
	`
	var client CreateClientIfNotExistResponse
	err := tx.QueryRow(ctx, query, req.IP, req.Count, req.Rate).
		Scan(&client.IP, &client.Count, &client.Rate)
	if err != nil {
		slog.Error("error inserting client to db", "error", err)
		return CreateClientIfNotExistResponse{}, fmt.Errorf("error add pvz in db: %w", err)
	}

	return client, nil
}
