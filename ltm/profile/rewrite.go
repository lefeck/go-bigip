package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

type RewriteList struct {
	Items    []Rewrite `json:"items,omitempty"`
	Kind     string    `json:"kind,omitempty"`
	SelfLink string    `json:"selflink,omitempty"`
}

type Rewrite struct {
	Kind                  string        `json:"kind"`
	Name                  string        `json:"name"`
	Partition             string        `json:"partition"`
	FullPath              string        `json:"fullPath"`
	Generation            int           `json:"generation"`
	SelfLink              string        `json:"selfLink"`
	AppService            string        `json:"appService"`
	BypassList            []interface{} `json:"bypassList"`
	ClientCachingType     string        `json:"clientCachingType"`
	DefaultsFrom          string        `json:"defaultsFrom"`
	DefaultsFromReference struct {
		Link string `json:"link"`
	} `json:"defaultsFromReference"`
	JavaCaFile          string `json:"javaCaFile"`
	JavaCaFileReference struct {
		Link string `json:"link"`
	} `json:"javaCaFileReference"`
	JavaCrl              string `json:"javaCrl"`
	JavaSignKey          string `json:"javaSignKey"`
	JavaSignKeyReference struct {
		Link string `json:"link"`
	} `json:"javaSignKeyReference"`
	JavaSigner          string `json:"javaSigner"`
	JavaSignerReference struct {
		Link string `json:"link"`
	} `json:"javaSignerReference"`
	LocationSpecific string `json:"locationSpecific"`
	Request          struct {
		InsertXforwardedFor   string `json:"insertXforwardedFor"`
		InsertXforwardedHost  string `json:"insertXforwardedHost"`
		InsertXforwardedProto string `json:"insertXforwardedProto"`
		RewriteHeaders        string `json:"rewriteHeaders"`
	} `json:"request"`
	Response struct {
		RewriteContent string `json:"rewriteContent"`
		RewriteHeaders string `json:"rewriteHeaders"`
	} `json:"response"`
	RewriteList       []interface{} `json:"rewriteList"`
	RewriteMode       string        `json:"rewriteMode"`
	SplitTunneling    string        `json:"splitTunneling"`
	URIRulesReference struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"uriRulesReference"`
}

const RewriteEndpoint = "rewrite"

type RewriteResource struct {
	b *bigip.BigIP
}

// List retrieves a list of Rewrite resources.
func (cr *RewriteResource) List() (*RewriteList, error) {
	var items RewriteList
	// Perform a GET request to retrieve a list of Rewrite resource objects
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RewriteEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into RewriteList struct
	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves a Rewrite resource by its full path name.
func (cr *RewriteResource) Get(fullPathName string) (*Rewrite, error) {
	var item Rewrite
	// Perform a GET request to retrieve a specific Rewrite resource by its full path name
	res, err := cr.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RewriteEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response data into Rewrite struct
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create adds a new Rewrite resource using the provided Rewrite item.
func (cr *RewriteResource) Create(item Rewrite) error {
	// Marshal the Rewrite struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a POST request to create a new Rewrite resource using the JSON data
	_, err = cr.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RewriteEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies a Rewrite resource identified by its full path name using the provided Rewrite item.
func (cr *RewriteResource) Update(fullPathName string, item Rewrite) error {
	// Marshal the Rewrite struct into JSON data
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)

	// Perform a PUT request to update the specified Rewrite resource with the JSON data
	_, err = cr.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RewriteEndpoint).SubResourceInstance(fullPathName).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a Rewrite resource by its full path name.
func (cr *RewriteResource) Delete(fullPathName string) error {
	// Perform a DELETE request to delete the specified Rewrite resource
	_, err := cr.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(LtmManager).
		Resource(ProfileEndpoint).SubResource(RewriteEndpoint).SubResourceInstance(fullPathName).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}
