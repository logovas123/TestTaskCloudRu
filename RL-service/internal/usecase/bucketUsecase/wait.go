package bucketUsecase

import (
	"context"
	"log/slog"

	"rl-service/internal/repository/bucketRepository"
)

type WaitRequest struct {
	IP string
}

// метод получает бакет из мапы по ip и удаляет токен из бакета
func (u *BucketUC) Wait(
	ctx context.Context,
	req WaitRequest,
) error {
	resp, err := u.bucketRepo.GetBucket(ctx, toGetBucketRequest(req))
	if err != nil {
		slog.Error("error getting bucket", "ip", req.IP, "error", err)
		return err
	}

	if err := resp.Bucket.Wait(ctx); err != nil {
		slog.Error("error remove token", "ip", req.IP, "error", err)
	}
	slog.Info("token success removed")

	return nil
}

func toGetBucketRequest(req WaitRequest) bucketRepository.GetBucketRequest {
	return bucketRepository.GetBucketRequest{
		IP: req.IP,
	}
}
