// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type DeleteGroupsByIDRequest struct {
	// Group ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteGroupsByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteGroupsByIDResponseBody - a list of ConfigGroup objects
type DeleteGroupsByIDResponseBody struct {
	// number of items present in the items array
	Count *int64                   `json:"count,omitempty"`
	Items []components.ConfigGroup `json:"items,omitempty"`
}

func (o *DeleteGroupsByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DeleteGroupsByIDResponseBody) GetItems() []components.ConfigGroup {
	if o == nil {
		return nil
	}
	return o.Items
}

type DeleteGroupsByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// a list of ConfigGroup objects
	Object *DeleteGroupsByIDResponseBody
}

func (o *DeleteGroupsByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteGroupsByIDResponse) GetObject() *DeleteGroupsByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
