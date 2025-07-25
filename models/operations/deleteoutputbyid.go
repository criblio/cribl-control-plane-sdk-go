// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

type DeleteOutputByIDRequest struct {
	// Unique ID to DELETE
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteOutputByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteOutputByIDResponseBody - a list of Destination objects
type DeleteOutputByIDResponseBody struct {
	// number of items present in the items array
	Count *int64              `json:"count,omitempty"`
	Items []components.Output `json:"items,omitempty"`
}

func (o *DeleteOutputByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DeleteOutputByIDResponseBody) GetItems() []components.Output {
	if o == nil {
		return nil
	}
	return o.Items
}

type DeleteOutputByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// a list of Destination objects
	Object *DeleteOutputByIDResponseBody
}

func (o *DeleteOutputByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteOutputByIDResponse) GetObject() *DeleteOutputByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
