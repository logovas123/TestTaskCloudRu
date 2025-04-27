package bucketUsecase

import (
	"context"

	"rl-service/internal/repository/bucketRepository"
)

type WaitRequest struct {
	IP string
}

func (u *BucketUC) Wait(
	ctx context.Context,
	req WaitRequest,
) error {
	resp, err := u.bucketRepo.GetBucket(ctx, toGetBucketRequest(req))
	if err != nil {
		return err
	}

	resp.Bucket.Wait(ctx)

	return nil
}

func toGetBucketRequest(req WaitRequest) bucketRepository.GetBucketRequest {
	return bucketRepository.GetBucketRequest{
		IP: req.IP,
	}
}
