package tfstate

import (
	"encoding/json"
	"strings"

	"github.com/mathematician/bifurcate/awsstate"
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

func GetAllResources(bucket string, keys []string) []Resource {
	allResources := []Resource{}

	for _, key := range keys {
		resources, err := GetResources(bucket, key)
		if err != nil {
			panic("error getting key, " + err.Error())
		}
		allResources = append(allResources, resources...)
	}

	return allResources
}

func GetResources(bucket string, key string) ([]Resource, error) {
	resources := []Resource{}
	stateBuf, err := awsstate.GetObject(bucket, key)
	if err != nil {
		return nil, err
	}

	state, err := extractStateData(stateBuf)

	for _, module := range state.Modules {
		if len(module.Resources) > 0 {
			for name, resource := range module.Resources {
				if strings.HasPrefix(name, "aws_") {
					resources = append(resources, Resource{
						Name: name,
						Type: resource.Type,
						ID:   resource.Primary.ID,
					})
				}
			}
		}
	}

	return resources, err

}

func extractStateData(byteData []byte) (State, error) {
	var state State
	err := json.Unmarshal(byteData, &state)

	return state, err
}
