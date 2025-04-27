package bucketRepository

import (
	"context"

	"rl-service/pkg/bucket"
	"rl-service/pkg/errlst"
)

type GetBucketRequest struct {
	IP string
}

type GetBucketResponse struct {
	Bucket *bucket.TokenBucket
}

// метод получает бакет по ip
func (r *BucketRepository) GetBucket(
	ctx context.Context,
	req GetBucketRequest,
) (GetBucketResponse, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	b, ok := r.listBuckets[req.IP]
	if !ok {
		return GetBucketResponse{}, errlst.ErrBucketNotExist
	}

	return GetBucketResponse{Bucket: b}, nil
}
