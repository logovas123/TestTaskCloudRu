package clientRepository

import "github.com/jackc/pgx/v5/pgxpool"

type ClientRepository struct {
	pool *pgxpool.Pool
}

func NewClientRepository(pool *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{
		pool: pool,
	}
}
