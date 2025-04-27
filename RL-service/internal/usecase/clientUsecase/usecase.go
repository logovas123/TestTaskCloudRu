package clientUsecase

import (
	"context"

	"rl-service/internal/config"
	"rl-service/internal/repository/clientRepository"
)

type ClientRepository interface {
	CreateClientIfNotExist(
		ctx context.Context,
		req clientRepository.CreateClientIfNotExistRequest,
	) (clientRepository.CreateClientIfNotExistResponse, error)
}

type ClientUC struct {
	cfg        *config.Config
	clientRepo ClientRepository
}

func NewClientUC(
	cfg *config.Config,
	clientRepo ClientRepository,
) *ClientUC {
	return &ClientUC{
		cfg:        cfg,
		clientRepo: clientRepo,
	}
}
