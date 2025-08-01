// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

type UpdateOutputByIDRequest struct {
	// Unique ID to PATCH
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Destination object to be updated
	Output components.Output `request:"mediaType=application/json"`
}

func (o *UpdateOutputByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateOutputByIDRequest) GetOutput() components.Output {
	if o == nil {
		return components.Output{}
	}
	return o.Output
}

// UpdateOutputByIDResponseBody - a list of Destination objects
type UpdateOutputByIDResponseBody struct {
	// number of items present in the items array
	Count *int64              `json:"count,omitempty"`
	Items []components.Output `json:"items,omitempty"`
}

func (o *UpdateOutputByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *UpdateOutputByIDResponseBody) GetItems() []components.Output {
	if o == nil {
		return nil
	}
	return o.Items
}

type UpdateOutputByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// a list of Destination objects
	Object *UpdateOutputByIDResponseBody
}

func (o *UpdateOutputByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateOutputByIDResponse) GetObject() *UpdateOutputByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
