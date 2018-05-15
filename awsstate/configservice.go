package awsstate

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/configserviceiface"
)

var services = []configservice.ResourceType{
	configservice.ResourceType("AWS::EC2::CustomerGateway"),
	configservice.ResourceType("AWS::EC2::EIP"),
	configservice.ResourceType("AWS::EC2::Host"),
	configservice.ResourceType("AWS::EC2::Instance"),
	configservice.ResourceType("AWS::EC2::InternetGateway"),
	configservice.ResourceType("AWS::EC2::NetworkAcl"),
	configservice.ResourceType("AWS::EC2::NetworkInterface"),
	configservice.ResourceType("AWS::EC2::RouteTable"),
	configservice.ResourceType("AWS::EC2::SecurityGroup"),
	configservice.ResourceType("AWS::EC2::Subnet"),
	configservice.ResourceType("AWS::CloudTrail::Trail"),
	configservice.ResourceType("AWS::EC2::Volume"),
	configservice.ResourceType("AWS::EC2::VPC"),
	configservice.ResourceType("AWS::EC2::VPNConnection"),
	configservice.ResourceType("AWS::EC2::VPNGateway"),
	configservice.ResourceType("AWS::IAM::Group"),
	configservice.ResourceType("AWS::IAM::Policy"),
	configservice.ResourceType("AWS::IAM::Role"),
	configservice.ResourceType("AWS::IAM::User"),
	configservice.ResourceType("AWS::ACM::Certificate"),
	configservice.ResourceType("AWS::RDS::DBInstance"),
	configservice.ResourceType("AWS::RDS::DBSubnetGroup"),
	configservice.ResourceType("AWS::RDS::DBSecurityGroup"),
	configservice.ResourceType("AWS::RDS::DBSnapshot"),
	configservice.ResourceType("AWS::RDS::EventSubscription"),
	configservice.ResourceType("AWS::ElasticLoadBalancingV2::LoadBalancer"),
	configservice.ResourceType("AWS::S3::Bucket"),
	configservice.ResourceType("AWS::SSM::ManagedInstanceInventory"),
	configservice.ResourceType("AWS::Redshift::Cluster"),
	configservice.ResourceType("AWS::Redshift::ClusterSnapshot"),
	configservice.ResourceType("AWS::Redshift::ClusterParameterGroup"),
	configservice.ResourceType("AWS::Redshift::ClusterSecurityGroup"),
	configservice.ResourceType("AWS::Redshift::ClusterSubnetGroup"),
	configservice.ResourceType("AWS::Redshift::EventSubscription"),
	configservice.ResourceType("AWS::CloudWatch::Alarm"),
	configservice.ResourceType("AWS::CloudFormation::Stack"),
	configservice.ResourceType("AWS::DynamoDB::Table"),
	configservice.ResourceType("AWS::AutoScaling::AutoScalingGroup"),
	configservice.ResourceType("AWS::AutoScaling::LaunchConfiguration"),
	configservice.ResourceType("AWS::AutoScaling::ScalingPolicy"),
	configservice.ResourceType("AWS::AutoScaling::ScheduledAction"),
	configservice.ResourceType("AWS::CodeBuild::Project"),
	configservice.ResourceType("AWS::WAF::RateBasedRule"),
	configservice.ResourceType("AWS::WAF::Rule"),
	configservice.ResourceType("AWS::WAF::WebACL"),
	configservice.ResourceType("AWS::WAFRegional::RateBasedRule"),
	configservice.ResourceType("AWS::WAFRegional::Rule"),
	configservice.ResourceType("AWS::WAFRegional::WebACL"),
	configservice.ResourceType("AWS::CloudFront::Distribution"),
	configservice.ResourceType("AWS::CloudFront::StreamingDistribution"),
}

func GetConfigServiceResourcesByService(configServiceClient configserviceiface.ConfigServiceAPI, service configservice.ResourceType) ([]configservice.ResourceIdentifier, error) {
	resources := []configservice.ResourceIdentifier{}
	params := &configservice.ListDiscoveredResourcesInput{
		ResourceType: service,
	}

	req := configServiceClient.ListDiscoveredResourcesRequest(params)
	resp, err := req.Send()

	for _, object := range resp.ResourceIdentifiers {
		resources = append(resources, object)
	}

	return resources, err
}

func GetAllConfigServiceResources(configServiceClient configserviceiface.ConfigServiceAPI, services []configservice.ResourceType) []configservice.ResourceIdentifier {
	allResources := []configservice.ResourceIdentifier{}

	for _, service := range services {
		resources, err := GetConfigServiceResourcesByService(configServiceClient, service)
		if err != nil {
			panic("error calling GetResourceByService, " + err.Error())
		}
		allResources = append(allResources, resources...)
	}

	return allResources
}

func GetConfigServiceResources() []configservice.ResourceIdentifier {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	configServiceClient := configservice.New(cfg)

	return GetAllConfigServiceResources(configServiceClient, services)
}

func main(service configservice.ResourceType) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	configServiceClient := configservice.New(cfg)

	GetConfigServiceResourcesByService(configServiceClient, service)
}
