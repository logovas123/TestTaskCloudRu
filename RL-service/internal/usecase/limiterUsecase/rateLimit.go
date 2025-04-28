package limiterUsecase

import (
	"context"
	"log/slog"

	"rl-service/internal/usecase/bucketUsecase"
	"rl-service/internal/usecase/clientUsecase"
)

/*
	На этом уровне (usecase) следующая логика:
	1) клиент проверяется на существование в базе; если клиент не существует, то в базу добавляется запись со
	следующими данными: ip, размер бакета, скорость добавления токенов в бакет
	2) После этого по параметру isExist проверяем, нужно ли стартовать новый бакет для клиента - если он не существовал,
	то стартуем новый
	3) После этого Wait вынимает токен из бакета и usecase заканчивает выполнение
*/

type RateLimitRequest struct {
	IP    string
	Count string
	Rate  string
}

func (u *LimiterUC) RateLimit(
	ctx context.Context,
	req RateLimitRequest,
) error {
	client, err := u.clientUC.CreateClientIfNotExist(ctx, toCreateClientIfNotExistRequest(req))
	if err != nil {
		return err
	}

	if !client.IsExist {
		slog.Info("client not existed, need create bucket", "ip", req.IP)
		if err := u.bucketUC.StartBucket(ctx, toStartBucketRequest(client)); err != nil {
			slog.Error("error start bucket", "ip", req.IP, "error", err)
			return err
		}
	}

	if err = u.bucketUC.Wait(ctx, toWaitRequest(client.IP)); err != nil {
		slog.Error("error remove token", "ip", req.IP, "error", err)
		return err
	}

	slog.Info("remove token success")

	return nil
}

func toCreateClientIfNotExistRequest(
	req RateLimitRequest,
) clientUsecase.CreateClientIfNotExistRequest {
	return clientUsecase.CreateClientIfNotExistRequest{
		IP:    req.IP,
		Count: req.Count,
		Rate:  req.Rate,
	}
}

func toStartBucketRequest(
	req clientUsecase.CreateClientIfNotExistResponse,
) bucketUsecase.StartBucketRequest {
	return bucketUsecase.StartBucketRequest{
		IP:        req.IP,
		Rate:      req.Rate,
		MaxTokens: req.Count,
	}
}

func toWaitRequest(ip string) bucketUsecase.WaitRequest {
	return bucketUsecase.WaitRequest{
		IP: ip,
	}
}
