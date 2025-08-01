// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type DeleteInputByIDRequest struct {
	// Unique ID to DELETE
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteInputByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteInputByIDResponseBody - a list of Source objects
type DeleteInputByIDResponseBody struct {
	// number of items present in the items array
	Count *int64             `json:"count,omitempty"`
	Items []components.Input `json:"items,omitempty"`
}

func (o *DeleteInputByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DeleteInputByIDResponseBody) GetItems() []components.Input {
	if o == nil {
		return nil
	}
	return o.Items
}

type DeleteInputByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// a list of Source objects
	Object *DeleteInputByIDResponseBody
}

func (o *DeleteInputByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteInputByIDResponse) GetObject() *DeleteInputByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
