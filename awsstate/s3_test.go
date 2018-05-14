package awsstate

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3/s3iface"
)

type mockS3Client struct {
	s3iface.S3API
}

func TestGetObjectBuffer(t *testing.T) {
}
