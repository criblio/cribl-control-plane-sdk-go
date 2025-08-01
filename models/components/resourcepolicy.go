// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ResourcePolicy struct {
	Gid    string       `json:"gid"`
	ID     *string      `json:"id,omitempty"`
	Policy string       `json:"policy"`
	Type   RbacResource `json:"type"`
}

func (o *ResourcePolicy) GetGid() string {
	if o == nil {
		return ""
	}
	return o.Gid
}

func (o *ResourcePolicy) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ResourcePolicy) GetPolicy() string {
	if o == nil {
		return ""
	}
	return o.Policy
}

func (o *ResourcePolicy) GetType() RbacResource {
	if o == nil {
		return RbacResource("")
	}
	return o.Type
}
