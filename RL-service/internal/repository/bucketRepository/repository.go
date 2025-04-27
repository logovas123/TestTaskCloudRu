package bucketRepository

import (
	"sync"

	"rl-service/pkg/bucket"
)

type BucketRepository struct {
	listBuckets map[string]*bucket.TokenBucket
	mutex       sync.Mutex
}

func NewBucketRepository() *BucketRepository {
	return &BucketRepository{
		listBuckets: map[string]*bucket.TokenBucket{},
		mutex:       sync.Mutex{},
	}
}
