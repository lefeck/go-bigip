// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package software

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lefeck/go-bigip"
	"strings"
)

// ImageList holds a list of Image configurations.
type ImageList struct {
	Items    []Image `json:"items"`
	Kind     string  `json:"kind"`
	SelfLink string  `json:"selflink"`
}

// Image holds the configuration of a single Image.
type Image struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	SelfLink     string `json:"selfLink"`
	Build        string `json:"build"`
	BuildDate    string `json:"buildDate"`
	Checksum     string `json:"checksum"`
	FileSize     string `json:"fileSize"`
	LastModified string `json:"lastModified"`
	Product      string `json:"product"`
	Verified     string `json:"verified"`
	Version      string `json:"version"`
}

// ImageEndpoint represents the REST resource for managing Image.
const ImageEndpoint = "image"

// ImageResource provides an API to manage Image configurations.
type ImageResource struct {
	b *bigip.BigIP
}

// List retrieves all Image details.
func (r *ImageResource) List() (*ImageList, error) {
	var items ImageList
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(ImageEndpoint).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(res, &items); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &items, nil
}

// Get retrieves the details of a single Image by node name.
func (r *ImageResource) Get(name string) (*Image, error) {
	var item Image
	res, err := r.b.RestClient.Get().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(ImageEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, &item); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %s\n", err)
	}
	return &item, nil
}

// Create creates a new Image item.
func (r *ImageResource) Create(item Image) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Post().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(ImageEndpoint).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Update modifies the Image item identified by the Image name.
func (r *ImageResource) Update(name string, item Image) error {
	jsonData, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %w", err)
	}
	jsonString := string(jsonData)
	_, err = r.b.RestClient.Put().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(ImageEndpoint).ResourceInstance(name).Body(strings.NewReader(jsonString)).DoRaw(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Delete a single Image identified by the Image name. if it is not exist return error
func (r *ImageResource) Delete(name string) error {
	_, err := r.b.RestClient.Delete().Prefix(bigip.GetBaseResource()).ResourceCategory(bigip.GetTMResource()).ManagerName(SysManager).
		Resource(SoftwareEndpoint).SubResource(ImageEndpoint).ResourceInstance(name).DoRaw(context.Background())
	if err != nil {
		return err
	}

	return nil
}
