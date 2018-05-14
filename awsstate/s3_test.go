package awsstate

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3/s3iface"
)

type mockS3Client struct {
	s3iface.S3API
}

func TestGetObjectBuffer(t *testing.T) {
	bucket := "awsdhubnp-terraform-state"
	key := "operations/bastion/terraform.tfstate"

	fmt.Printf("Get object at: %s/%s.", bucket, key)
	//_, err := getObjectBuffer(&mockS3Client{}, bucket, key)
}
