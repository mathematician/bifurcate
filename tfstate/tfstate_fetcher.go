package tfstate

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
)

func readStateFromS3(s3Client *s3.S3, bucket string) (*State, error) {

	buf, err := downloadS3Data(s3Client, bucket, key)
	if err != nil {
		return nil, err
	}

	return buf
}

func downloadS3Data(s3Client *s3.S3, bucket string, key string) ([]byte, error) {
	buf := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloaderWithClient(s3Client)

	_, err := downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
