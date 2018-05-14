package awsstate

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/configserviceiface"
)

type mockConfigServiceClient struct {
	configserviceiface.ConfigServiceAPI
}

func (m *mockConfigServiceClient) ListDiscoveredResourcesRequest(input *configservice.ListDiscoveredResourcesInput) configservice.ListDiscoveredResourcesRequest {
	op := &aws.Operation{
		Name:       "ListDiscoveredResources",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &configservice.ListDiscoveredResourcesInput{}
	}

	output := &configservice.ListDiscoveredResourcesOutput{}
	req := m.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return configservice.ListDiscoveredResourcesRequest{Request: req, Input: input, Copy: m.ListDiscoveredResourcesRequest}
	//	return &configservice.ListDiscoveredResourcesRequest{}
}

func TestGetResourcesByService(t *testing.T) {
	service := configservice.ResourceType("AWS::EC2::Instance")

	mockConfigServiceClient := &mockConfigServiceClient{}

	resources, err := GetResourcesByService(mockConfigServiceClient, service)
	if err != nil {
		panic("failed, " + err.Error())
	}

	fmt.Printf("resources: %s", resources)
}
