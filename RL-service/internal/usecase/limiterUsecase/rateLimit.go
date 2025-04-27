package limiterUsecase

import (
	"context"

	"rl-service/internal/usecase/bucketUsecase"
	"rl-service/internal/usecase/clientUsecase"
)

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
		if err := u.bucketUC.StartBucket(ctx, toStartBucketRequest(client)); err != nil {
			return err
		}
	}

	if err = u.bucketUC.Wait(ctx, toWaitRequest(client.IP)); err != nil {
		return err
	}

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
