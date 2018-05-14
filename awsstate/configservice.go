package awsstate

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/configserviceiface"
)

func GetResourcesByService(configServiceClient configserviceiface.ConfigServiceAPI, service configservice.ResourceType) ([]configservice.ResourceIdentifier, error) {
	resources := []configservice.ResourceIdentifier{}
	params := &configservice.ListDiscoveredResourcesInput{
		ResourceType: service,
	}

	req := configServiceClient.ListDiscoveredResourcesRequest(params)
	resp, err := req.Send()

	for _, object := range resp.ResourceIdentifiers {
		resources = append(resources, object)
	}

	fmt.Printf("resources: %s", resources)
	return resources, err
}

func main(service configservice.ResourceType) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	configServiceClient := configservice.New(cfg)

	GetResourcesByService(configServiceClient, service)
}
