package bucketUsecase

import (
	"context"

	"rl-service/internal/config"
	"rl-service/internal/repository/bucketRepository"
)

type BucketRepository interface {
	AddBucket(
		ctx context.Context,
		req bucketRepository.AddBucketRequest,
	) error
	GetBucket(
		ctx context.Context,
		req bucketRepository.GetBucketRequest,
	) (bucketRepository.GetBucketResponse, error)
}

type BucketUC struct {
	cfg        *config.Config
	bucketRepo BucketRepository
}

func NewBucketUC(
	cfg *config.Config,
	bucketRepo BucketRepository,
) *BucketUC {
	return &BucketUC{
		cfg:        cfg,
		bucketRepo: bucketRepo,
	}
}
