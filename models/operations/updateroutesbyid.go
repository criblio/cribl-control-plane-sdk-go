// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

type UpdateRoutesByIDRequest struct {
	// Unique ID to PATCH
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Routes object to be updated
	Routes components.RoutesInput `request:"mediaType=application/json"`
}

func (o *UpdateRoutesByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateRoutesByIDRequest) GetRoutes() components.RoutesInput {
	if o == nil {
		return components.RoutesInput{}
	}
	return o.Routes
}

// UpdateRoutesByIDResponseBody - a list of Routes objects
type UpdateRoutesByIDResponseBody struct {
	// number of items present in the items array
	Count *int64              `json:"count,omitempty"`
	Items []components.Routes `json:"items,omitempty"`
}

func (o *UpdateRoutesByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *UpdateRoutesByIDResponseBody) GetItems() []components.Routes {
	if o == nil {
		return nil
	}
	return o.Items
}

type UpdateRoutesByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// a list of Routes objects
	Object *UpdateRoutesByIDResponseBody
}

func (o *UpdateRoutesByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateRoutesByIDResponse) GetObject() *UpdateRoutesByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
