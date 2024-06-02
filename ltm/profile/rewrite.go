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
	Kind                  string        `json:"kind,omitempty"`
	Name                  string        `json:"name,omitempty"`
	Partition             string        `json:"partition,omitempty"`
	FullPath              string        `json:"fullPath,omitempty"`
	Generation            int           `json:"generation,omitempty"`
	SelfLink              string        `json:"selfLink,omitempty"`
	AppService            string        `json:"appService,omitempty"`
	BypassList            []interface{} `json:"bypassList,omitempty"`
	ClientCachingType     string        `json:"clientCachingType,omitempty"`
	DefaultsFrom          string        `json:"defaultsFrom,omitempty"`
	DefaultsFromReference struct {
		Link string `json:"link,omitempty"`
	} `json:"defaultsFromReference,omitempty"`
	JavaCaFile          string `json:"javaCaFile,omitempty"`
	JavaCaFileReference struct {
		Link string `json:"link,omitempty"`
	} `json:"javaCaFileReference,omitempty"`
	JavaCrl              string `json:"javaCrl,omitempty"`
	JavaSignKey          string `json:"javaSignKey,omitempty"`
	JavaSignKeyReference struct {
		Link string `json:"link,omitempty"`
	} `json:"javaSignKeyReference,omitempty"`
	JavaSigner          string `json:"javaSigner,omitempty"`
	JavaSignerReference struct {
		Link string `json:"link,omitempty"`
	} `json:"javaSignerReference,omitempty"`
	LocationSpecific string `json:"locationSpecific,omitempty"`
	Request          struct {
		InsertXforwardedFor   string `json:"insertXforwardedFor,omitempty"`
		InsertXforwardedHost  string `json:"insertXforwardedHost,omitempty"`
		InsertXforwardedProto string `json:"insertXforwardedProto,omitempty"`
		RewriteHeaders        string `json:"rewriteHeaders,omitempty"`
	} `json:"request,omitempty"`
	Response struct {
		RewriteContent string `json:"rewriteContent,omitempty"`
		RewriteHeaders string `json:"rewriteHeaders,omitempty"`
	} `json:"response,omitempty"`
	RewriteList       []interface{} `json:"rewriteList,omitempty"`
	RewriteMode       string        `json:"rewriteMode,omitempty"`
	SplitTunneling    string        `json:"splitTunneling,omitempty"`
	URIRulesReference struct {
		Link            string `json:"link,omitempty"`
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
	} `json:"uriRulesReference,omitempty"`
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
