package tfstate

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
)

type State struct {
	Modules []*moduleState `json:"modules"`
}

type Resource struct {
	Name string
	Type string
	ID   string
}

type moduleState struct {
	Resources map[string]*resourceState `json:"resources"`
}

type resourceState struct {
	Type    string         `json:"type"`
	Primary *instanceState `json:"primary"`
}

type instanceState struct {
	ID         string            `json:"id"`
	Attributes map[string]string `json:"attributes"`
}

func GetResources(bucket string) ([]Resource, error) {
	resources := []Resource{}
	state, err := GetState(bucket)
	if err != nil {
		return nil, err
	}

	for _, module := range state.Modules {
		if len(module.Resources) > 0 {
			for name, resource := range module.Resources {
				if !strings.HasPrefix(name, "data.") {
					resources = append(resources, Resource{
						Name: name,
						Type: resource.Type,
						ID:   resource.Primary.ID,
					})
				}
			}
		}
	}

	return resources, nil

}

func GetState(bucket string) (*State, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	s3Client := s3.New(cfg)

	return readStateFromS3(s3Client, bucket)
}

func readStateFromS3(s3Client *s3.S3, bucket string) (*State, error) {
	key := "operations/bastion/terraform.tfstate"

	buf, err := downloadS3Data(s3Client, bucket, key)
	if err != nil {
		return nil, err
	}

	return extractStateData(buf)
}

func extractStateData(byteData []byte) (*State, error) {
	var state *State
	err := json.Unmarshal(byteData, &state)
	if err != nil {
		panic("Error, " + err.Error())
		return nil, err
	}

	return state, nil
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
