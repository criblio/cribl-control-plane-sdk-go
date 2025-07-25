// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

// CreatePipelineResponseBody - a list of Pipeline objects
type CreatePipelineResponseBody struct {
	// number of items present in the items array
	Count *int64                `json:"count,omitempty"`
	Items []components.Pipeline `json:"items,omitempty"`
}

func (o *CreatePipelineResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *CreatePipelineResponseBody) GetItems() []components.Pipeline {
	if o == nil {
		return nil
	}
	return o.Items
}

type CreatePipelineResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// a list of Pipeline objects
	Object *CreatePipelineResponseBody
}

func (o *CreatePipelineResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreatePipelineResponse) GetObject() *CreatePipelineResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
