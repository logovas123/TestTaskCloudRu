package bucketUsecase

import (
	"context"
	"log/slog"

	"rl-service/internal/repository/bucketRepository"
	"rl-service/pkg/bucket"
)

type StartBucketRequest struct {
	IP        string
	Rate      uint64
	MaxTokens uint64
}

// в методе создаётся бакет, далее бакет добавляется в мапу и стартуется
func (u *BucketUC) StartBucket(
	ctx context.Context,
	req StartBucketRequest,
) error {
	bkt := bucket.NewTokenBucket(req.MaxTokens, req.Rate)

	err := u.bucketRepo.AddBucket(ctx, toAddBucketRequest(req, bkt))
	if err != nil {
		slog.Error("error add bucket to map", "ip", req.IP, "error", err)
		return err
	}

	bkt.Start(context.Background())

	slog.Info("bucket start success")

	return nil
}

func toAddBucketRequest(req StartBucketRequest, b *bucket.TokenBucket) bucketRepository.AddBucketRequest {
	return bucketRepository.AddBucketRequest{
		IP:     req.IP,
		Bucket: b,
	}
}
