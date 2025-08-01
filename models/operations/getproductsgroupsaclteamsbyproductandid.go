// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

// GetProductsGroupsACLTeamsByProductAndIDProduct - Cribl Product
type GetProductsGroupsACLTeamsByProductAndIDProduct string

const (
	GetProductsGroupsACLTeamsByProductAndIDProductStream GetProductsGroupsACLTeamsByProductAndIDProduct = "stream"
	GetProductsGroupsACLTeamsByProductAndIDProductEdge   GetProductsGroupsACLTeamsByProductAndIDProduct = "edge"
)

func (e GetProductsGroupsACLTeamsByProductAndIDProduct) ToPointer() *GetProductsGroupsACLTeamsByProductAndIDProduct {
	return &e
}
func (e *GetProductsGroupsACLTeamsByProductAndIDProduct) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stream":
		fallthrough
	case "edge":
		*e = GetProductsGroupsACLTeamsByProductAndIDProduct(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProductsGroupsACLTeamsByProductAndIDProduct: %v", v)
	}
}

// GetProductsGroupsACLTeamsByProductAndIDType - resource type by which to filter access levels
type GetProductsGroupsACLTeamsByProductAndIDType string

const (
	GetProductsGroupsACLTeamsByProductAndIDTypeGroups           GetProductsGroupsACLTeamsByProductAndIDType = "groups"
	GetProductsGroupsACLTeamsByProductAndIDTypeDatasets         GetProductsGroupsACLTeamsByProductAndIDType = "datasets"
	GetProductsGroupsACLTeamsByProductAndIDTypeDatasetProviders GetProductsGroupsACLTeamsByProductAndIDType = "dataset-providers"
	GetProductsGroupsACLTeamsByProductAndIDTypeProjects         GetProductsGroupsACLTeamsByProductAndIDType = "projects"
	GetProductsGroupsACLTeamsByProductAndIDTypeDashboards       GetProductsGroupsACLTeamsByProductAndIDType = "dashboards"
	GetProductsGroupsACLTeamsByProductAndIDTypeMacros           GetProductsGroupsACLTeamsByProductAndIDType = "macros"
	GetProductsGroupsACLTeamsByProductAndIDTypeNotebooks        GetProductsGroupsACLTeamsByProductAndIDType = "notebooks"
	GetProductsGroupsACLTeamsByProductAndIDTypeInsights         GetProductsGroupsACLTeamsByProductAndIDType = "insights"
)

func (e GetProductsGroupsACLTeamsByProductAndIDType) ToPointer() *GetProductsGroupsACLTeamsByProductAndIDType {
	return &e
}
func (e *GetProductsGroupsACLTeamsByProductAndIDType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "groups":
		fallthrough
	case "datasets":
		fallthrough
	case "dataset-providers":
		fallthrough
	case "projects":
		fallthrough
	case "dashboards":
		fallthrough
	case "macros":
		fallthrough
	case "notebooks":
		fallthrough
	case "insights":
		*e = GetProductsGroupsACLTeamsByProductAndIDType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProductsGroupsACLTeamsByProductAndIDType: %v", v)
	}
}

type GetProductsGroupsACLTeamsByProductAndIDRequest struct {
	// Cribl Product
	Product GetProductsGroupsACLTeamsByProductAndIDProduct `pathParam:"style=simple,explode=false,name=product"`
	// Group ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// resource type by which to filter access levels
	Type *GetProductsGroupsACLTeamsByProductAndIDType `queryParam:"style=form,explode=true,name=type"`
}

func (o *GetProductsGroupsACLTeamsByProductAndIDRequest) GetProduct() GetProductsGroupsACLTeamsByProductAndIDProduct {
	if o == nil {
		return GetProductsGroupsACLTeamsByProductAndIDProduct("")
	}
	return o.Product
}

func (o *GetProductsGroupsACLTeamsByProductAndIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetProductsGroupsACLTeamsByProductAndIDRequest) GetType() *GetProductsGroupsACLTeamsByProductAndIDType {
	if o == nil {
		return nil
	}
	return o.Type
}

// GetProductsGroupsACLTeamsByProductAndIDResponseBody - a list of TeamAccessControlList objects
type GetProductsGroupsACLTeamsByProductAndIDResponseBody struct {
	// number of items present in the items array
	Count *int64                             `json:"count,omitempty"`
	Items []components.TeamAccessControlList `json:"items,omitempty"`
}

func (o *GetProductsGroupsACLTeamsByProductAndIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetProductsGroupsACLTeamsByProductAndIDResponseBody) GetItems() []components.TeamAccessControlList {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetProductsGroupsACLTeamsByProductAndIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// a list of TeamAccessControlList objects
	Object *GetProductsGroupsACLTeamsByProductAndIDResponseBody
}

func (o *GetProductsGroupsACLTeamsByProductAndIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetProductsGroupsACLTeamsByProductAndIDResponse) GetObject() *GetProductsGroupsACLTeamsByProductAndIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
