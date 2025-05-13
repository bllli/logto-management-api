package components

import (
	"github.com/bllli/logto-management-api/internal/utils"
)

type Security struct {
	ClientID     *string `security:"scheme,type=oauth2,subtype=client_credentials,name=clientID,env=logtomanagementapi_client_id"`
	ClientSecret *string `security:"scheme,type=oauth2,subtype=client_credentials,name=clientSecret,env=logtomanagementapi_client_secret"`
	TokenURL     *string `default:"/oidc/token"`
	Resource     *string `security:"scheme,type=oauth2,subtype=client_credentials,name=resource,env=logtomanagementapi_resource"`
	Scope        *string `security:"scheme,type=oauth2,subtype=client_credentials,name=scope,env=logtomanagementapi_scope"`
}

func (s Security) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Security) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Security) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *Security) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *Security) GetTokenURL() *string {
	if o == nil {
		return nil
	}
	return o.TokenURL
}

func (o *Security) GetResource() *string {
	if o == nil {
		return nil
	}
	return o.Resource
}

func (o *Security) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}
