package s3iface

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3"
)

// Service reads and writes to the given bucket.
type Service interface {
	Put(string, io.Reader) (*s3.PutObjectOutput, error)
	Read(string) (*io.ReadCloser, error)
	ReadUnmarshal(string, interface{}) error
	SetACL(*string)
	SetSSE(*string)
	ListObjectsKeysV2Pages(*s3.ListObjectsV2Input) ([]string, bool, error)
	ListObjectsV2Input() *s3.ListObjectsV2Input
}
