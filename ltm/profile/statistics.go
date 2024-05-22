package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type StatisticsList struct {
	Items    []Statistics `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

type Statistics struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	AppService   string `json:"appService"`
	DefaultsFrom string `json:"defaultsFrom"`
	Description  string `json:"description"`
	Field1       string `json:"field1"`
	Field10      string `json:"field10"`
	Field11      string `json:"field11"`
	Field12      string `json:"field12"`
	Field13      string `json:"field13"`
	Field14      string `json:"field14"`
	Field15      string `json:"field15"`
	Field16      string `json:"field16"`
	Field17      string `json:"field17"`
	Field18      string `json:"field18"`
	Field19      string `json:"field19"`
	Field2       string `json:"field2"`
	Field20      string `json:"field20"`
	Field21      string `json:"field21"`
	Field22      string `json:"field22"`
	Field23      string `json:"field23"`
	Field24      string `json:"field24"`
	Field25      string `json:"field25"`
	Field26      string `json:"field26"`
	Field27      string `json:"field27"`
	Field28      string `json:"field28"`
	Field29      string `json:"field29"`
	Field3       string `json:"field3"`
	Field30      string `json:"field30"`
	Field31      string `json:"field31"`
	Field32      string `json:"field32"`
	Field4       string `json:"field4"`
	Field5       string `json:"field5"`
	Field6       string `json:"field6"`
	Field7       string `json:"field7"`
	Field8       string `json:"field8"`
	Field9       string `json:"field9"`
}

const StatisticsEndpoint = "stats"

type StatisticsResource struct {
	b *bigip.BigIP
}

// List retrieves a list of Statistics resources.
func (cr *StatisticsResource) List() (*StatisticsList, error) {
	var items StatisticsList
	// Perform a GET request to retrieve a list of Statistics resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(StatisticsEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into StatisticsList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a Statistics resource by its full path name.
func (cr *StatisticsResource) Get(fullPathName string) (*Statistics, error) {
	var item Statistics
	// Perform a GET request to retrieve a specific Statistics resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(StatisticsEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into Statistics struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new Statistics resource using the provided Statistics item.
func (cr *StatisticsResource) Create(item Statistics) error {
	// Marshal the Statistics struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new Statistics resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(StatisticsEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a Statistics resource identified by its full path name using the provided Statistics item.
func (cr *StatisticsResource) Update(fullPathName string, item Statistics) error {
	// Marshal the Statistics struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified Statistics resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(StatisticsEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a Statistics resource by its full path name.
func (cr *StatisticsResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified Statistics resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(StatisticsEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
