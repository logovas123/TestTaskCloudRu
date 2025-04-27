package limiterUsecase

import (
	"context"

	"rl-service/internal/config"
	"rl-service/internal/usecase/bucketUsecase"
	"rl-service/internal/usecase/clientUsecase"
)

type (
	ClientUC interface {
		CreateClientIfNotExist(
			ctx context.Context,
			req clientUsecase.CreateClientIfNotExistRequest,
		) (clientUsecase.CreateClientIfNotExistResponse, error)
	}

	BucketUC interface {
		StartBucket(
			ctx context.Context,
			req bucketUsecase.StartBucketRequest,
		) error
		Wait(
			ctx context.Context,
			req bucketUsecase.WaitRequest,
		) error
	}
)

type LimiterUC struct {
	cfg      *config.Config
	clientUC ClientUC
	bucketUC BucketUC
}

func NewLimiterUC(
	cfg *config.Config,
	clientUC ClientUC,
	bucketUC BucketUC,
) *LimiterUC {
	return &LimiterUC{
		cfg:      cfg,
		clientUC: clientUC,
		bucketUC: bucketUC,
	}
}
