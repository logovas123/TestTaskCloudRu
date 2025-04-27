package bucketRepository

import (
	"context"

	"rl-service/pkg/bucket"
	"rl-service/pkg/errlst"
)

type AddBucketRequest struct {
	IP     string
	Bucket *bucket.TokenBucket
}

func (r *BucketRepository) AddBucket(
	ctx context.Context,
	req AddBucketRequest,
) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, ok := r.listBuckets[req.IP]; ok {
		return errlst.ErrBucketExist
	}
	r.listBuckets[req.IP] = req.Bucket
	return nil
}
