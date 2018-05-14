package awsstate

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3iface"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
)

func FindKeysBySuffix(bucket string, suffixFilter string) ([]string, error) {
	keys := []string{}
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	s3Client := s3.New(cfg)

	resp, err := FindKeys(s3Client, bucket)

	for _, object := range resp.Contents {
		if strings.HasSuffix(*object.Key, suffixFilter) {
			keys = append(keys, *object.Key)
		}
	}

	return keys, err
}

func FindKeys(s3Client s3iface.S3API, bucket string) (*s3.ListObjectsV2Output, error) {
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}

	req := s3Client.ListObjectsV2Request(params)

	return req.Send()
}

func GetObject(bucket string, key string) ([]byte, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	s3Client := s3.New(cfg)

	buf, err := getObjectBuffer(s3Client, bucket, key)
	if err != nil {
		return nil, err
	}

	return buf, err
}

func getObjectBuffer(s3Client s3iface.S3API, bucket string, key string) ([]byte, error) {
	buf := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloaderWithClient(s3Client)

	_, err := downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	return buf.Bytes(), err
}
