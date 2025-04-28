package clientUsecase

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"rl-service/internal/repository/clientRepository"
	"rl-service/pkg/errlst"
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

// метод для создания клиента
func (u *ClientUC) CreateClientIfNotExist(
	ctx context.Context,
	req CreateClientIfNotExistRequest,
) (CreateClientIfNotExistResponse, error) {
	reqNew, err := toCreateClientIfNotExistRequest(req)
	if err != nil {
		slog.Error("error convert request", "ip", req.IP, "error", err)
		return CreateClientIfNotExistResponse{}, fmt.Errorf("error convert request: %w", err)
	}
	client, err := u.clientRepo.CreateClientIfNotExist(ctx, reqNew)
	if err != nil {
		return CreateClientIfNotExistResponse{}, fmt.Errorf("error creating client: %w", err)
	}

	return toCreateClientIfNotExistResponse(client), nil
}

func toCreateClientIfNotExistRequest(req CreateClientIfNotExistRequest) (
	clientRepository.CreateClientIfNotExistRequest,
	error,
) {
	count, err := strconv.ParseUint(req.Count, 10, 64)
	if err != nil {
		return clientRepository.CreateClientIfNotExistRequest{}, fmt.Errorf("invalid count: %w", err)
	}
	if count == 0 {
		return clientRepository.CreateClientIfNotExistRequest{},
			errlst.ErrCountOrRateNoValid
	}

	rate, err := strconv.ParseUint(req.Rate, 10, 64)
	if err != nil {
		return clientRepository.CreateClientIfNotExistRequest{}, fmt.Errorf("invalid rate: %w", err)
	}
	if rate == 0 {
		return clientRepository.CreateClientIfNotExistRequest{},
			errlst.ErrCountOrRateNoValid
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
