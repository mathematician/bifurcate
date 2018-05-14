package awsstate

import (
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

	output := &configservice.ListDiscoveredResourcesOutput{
		ResourceIdentifiers: []configservice.ResourceIdentifier{
			configservice.ResourceIdentifier{
				ResourceId:   aws.String("i-11111111111111111"),
				ResourceType: "AWS::EC2::Instance",
			},
			configservice.ResourceIdentifier{
				ResourceId:   aws.String("i-22222222222222222"),
				ResourceType: "AWS::EC2::Instance",
			},
		},
	}
	req := &aws.Request{
		Operation: op,
		Params:    input,
		Data:      output,
	}

	return configservice.ListDiscoveredResourcesRequest{Request: req, Input: input, Copy: m.ListDiscoveredResourcesRequest}
}

func TestGetResourcesByService(t *testing.T) {
	service := configservice.ResourceType("AWS::EC2::Instance")

	mockConfigServiceClient := &mockConfigServiceClient{}

	resources, err := GetResourcesByService(mockConfigServiceClient, service)
	if err != nil {
		panic("failed, " + err.Error())
	}

	if *resources[0].ResourceId != "i-11111111111111111" || *resources[1].ResourceId != "i-22222222222222222" {
		t.Errorf("Values not correct, got: %s and %s, want: %s and %s.", *resources[0].ResourceId, *resources[1].ResourceId, "i-11111111111111111", "i-22222222222222222")
	}
}
