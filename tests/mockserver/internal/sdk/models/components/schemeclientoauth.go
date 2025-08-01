// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
)

type SchemeClientOauth struct {
	ClientID     string `security:"name=clientID"`
	ClientSecret string `security:"name=clientSecret"`
	TokenURL     string `default:"https://login.cribl.cloud/oauth/token"`
	Audience     string `security:"name=audience"`
}

func (s SchemeClientOauth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SchemeClientOauth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SchemeClientOauth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SchemeClientOauth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SchemeClientOauth) GetTokenURL() string {
	if o == nil {
		return ""
	}
	return o.TokenURL
}

func (o *SchemeClientOauth) GetAudience() string {
	if o == nil {
		return ""
	}
	return o.Audience
}
