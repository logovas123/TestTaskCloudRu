package errlst

import "errors"

var (
	ErrBucketExist    = errors.New("bucket exist")
	ErrBucketNotExist = errors.New("bucket not exist")
)
