package bucketUsecase

import (
	"context"

	"rl-service/internal/repository/bucketRepository"
	"rl-service/pkg/bucket"
)

type StartBucketRequest struct {
	IP        string
	Rate      uint64
	MaxTokens uint64
}

func (u *BucketUC) StartBucket(
	ctx context.Context,
	req StartBucketRequest,
) error {
	bucket := bucket.NewTokenBucket(req.MaxTokens, req.Rate)

	err := u.bucketRepo.AddBucket(ctx, toAddBucketRequest(req, bucket))
	if err != nil {
		return err
	}

	bucket.Start(ctx)

	return nil
}

func toAddBucketRequest(req StartBucketRequest, b *bucket.TokenBucket) bucketRepository.AddBucketRequest {
	return bucketRepository.AddBucketRequest{
		IP:     req.IP,
		Bucket: b,
	}
}
