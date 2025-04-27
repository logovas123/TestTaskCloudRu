package clientUsecase

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"rl-service/internal/repository/clientRepository"
)

type CreateClientIfNotExistRequest struct {
	IP    string
	Count string
	Rate  string
}

type CreateClientIfNotExistResponse struct {
	IP      string
	Count   uint64
	Rate    uint64
	IsExist bool
}

func (u *ClientUC) CreateClientIfNotExist(
	ctx context.Context,
	req CreateClientIfNotExistRequest,
) (CreateClientIfNotExistResponse, error) {
	reqNew, err := toCreateClientIfNotExistRequest(req)
	if err != nil {
		slog.Error("error convert request", "error", err)
		return CreateClientIfNotExistResponse{}, fmt.Errorf("error convert request: %w", err)
	}
	client, err := u.clientRepo.CreateClientIfNotExist(ctx, reqNew)
	if err != nil {
		slog.Error("error creating client", "error", err)
		return CreateClientIfNotExistResponse{}, fmt.Errorf("error creating client: %w", err)
	}

	return toCreateClientIfNotExistResponse(client), nil
}

func toCreateClientIfNotExistRequest(req CreateClientIfNotExistRequest) (
	clientRepository.CreateClientIfNotExistRequest,
	error,
) {
	count, err := strconv.Atoi(req.Count)
	if err != nil {
		return clientRepository.CreateClientIfNotExistRequest{}, fmt.Errorf("invalid count: %w", err)
	}

	rate, err := strconv.Atoi(req.Rate)
	if err != nil {
		return clientRepository.CreateClientIfNotExistRequest{}, fmt.Errorf("invalid rate: %w", err)
	}

	return clientRepository.CreateClientIfNotExistRequest{
		IP:    req.IP,
		Count: uint64(count),
		Rate:  uint64(rate),
	}, nil
}

func toCreateClientIfNotExistResponse(
	req clientRepository.CreateClientIfNotExistResponse,
) CreateClientIfNotExistResponse {
	return CreateClientIfNotExistResponse{
		IP:      req.IP,
		Count:   req.Count,
		Rate:    req.Rate,
		IsExist: req.IsExist,
	}
}
