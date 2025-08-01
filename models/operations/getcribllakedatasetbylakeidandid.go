// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

type GetCriblLakeDatasetByLakeIDAndIDRequest struct {
	// lake id that contains the Datasets
	LakeID string `pathParam:"style=simple,explode=false,name=lakeId"`
	// dataset id to get
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCriblLakeDatasetByLakeIDAndIDRequest) GetLakeID() string {
	if o == nil {
		return ""
	}
	return o.LakeID
}

func (o *GetCriblLakeDatasetByLakeIDAndIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetCriblLakeDatasetByLakeIDAndIDResponseBody - a list of CriblLakeDataset objects
type GetCriblLakeDatasetByLakeIDAndIDResponseBody struct {
	// number of items present in the items array
	Count *int64                        `json:"count,omitempty"`
	Items []components.CriblLakeDataset `json:"items,omitempty"`
}

func (o *GetCriblLakeDatasetByLakeIDAndIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetCriblLakeDatasetByLakeIDAndIDResponseBody) GetItems() []components.CriblLakeDataset {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetCriblLakeDatasetByLakeIDAndIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// a list of CriblLakeDataset objects
	Object *GetCriblLakeDatasetByLakeIDAndIDResponseBody
}

func (o *GetCriblLakeDatasetByLakeIDAndIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCriblLakeDatasetByLakeIDAndIDResponse) GetObject() *GetCriblLakeDatasetByLakeIDAndIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
