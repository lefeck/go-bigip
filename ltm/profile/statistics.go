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
	Kind         string `json:"kind,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	AppService   string `json:"appService,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Field1       string `json:"field1,omitempty"`
	Field10      string `json:"field10,omitempty"`
	Field11      string `json:"field11,omitempty"`
	Field12      string `json:"field12,omitempty"`
	Field13      string `json:"field13,omitempty"`
	Field14      string `json:"field14,omitempty"`
	Field15      string `json:"field15,omitempty"`
	Field16      string `json:"field16,omitempty"`
	Field17      string `json:"field17,omitempty"`
	Field18      string `json:"field18,omitempty"`
	Field19      string `json:"field19,omitempty"`
	Field2       string `json:"field2,omitempty"`
	Field20      string `json:"field20,omitempty"`
	Field21      string `json:"field21,omitempty"`
	Field22      string `json:"field22,omitempty"`
	Field23      string `json:"field23,omitempty"`
	Field24      string `json:"field24,omitempty"`
	Field25      string `json:"field25,omitempty"`
	Field26      string `json:"field26,omitempty"`
	Field27      string `json:"field27,omitempty"`
	Field28      string `json:"field28,omitempty"`
	Field29      string `json:"field29,omitempty"`
	Field3       string `json:"field3,omitempty"`
	Field30      string `json:"field30,omitempty"`
	Field31      string `json:"field31,omitempty"`
	Field32      string `json:"field32,omitempty"`
	Field4       string `json:"field4,omitempty"`
	Field5       string `json:"field5,omitempty"`
	Field6       string `json:"field6,omitempty"`
	Field7       string `json:"field7,omitempty"`
	Field8       string `json:"field8,omitempty"`
	Field9       string `json:"field9,omitempty"`
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
