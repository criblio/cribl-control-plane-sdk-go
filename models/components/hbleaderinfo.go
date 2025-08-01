// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type HBLeaderInfo struct {
	Host       string  `json:"host"`
	Port       float64 `json:"port"`
	Servername *string `json:"servername,omitempty"`
	TLS        *bool   `json:"tls,omitempty"`
}

func (o *HBLeaderInfo) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *HBLeaderInfo) GetPort() float64 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

func (o *HBLeaderInfo) GetServername() *string {
	if o == nil {
		return nil
	}
	return o.Servername
}

func (o *HBLeaderInfo) GetTLS() *bool {
	if o == nil {
		return nil
	}
	return o.TLS
}
