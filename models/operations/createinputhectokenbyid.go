// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

type CreateInputHecTokenByIDRequest struct {
	// HEC Source id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// AddHecTokenRequest object
	AddHecTokenRequest components.AddHecTokenRequest `request:"mediaType=application/json"`
}

func (o *CreateInputHecTokenByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateInputHecTokenByIDRequest) GetAddHecTokenRequest() components.AddHecTokenRequest {
	if o == nil {
		return components.AddHecTokenRequest{}
	}
	return o.AddHecTokenRequest
}

// CreateInputHecTokenByIDResponseBody - a list of any objects
type CreateInputHecTokenByIDResponseBody struct {
	// number of items present in the items array
	Count *int64           `json:"count,omitempty"`
	Items []map[string]any `json:"items,omitempty"`
}

func (o *CreateInputHecTokenByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *CreateInputHecTokenByIDResponseBody) GetItems() []map[string]any {
	if o == nil {
		return nil
	}
	return o.Items
}

type CreateInputHecTokenByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// a list of any objects
	Object *CreateInputHecTokenByIDResponseBody
}

func (o *CreateInputHecTokenByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateInputHecTokenByIDResponse) GetObject() *CreateInputHecTokenByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
