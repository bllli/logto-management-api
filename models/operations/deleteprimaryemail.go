// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/bllli/logto-management-api/models/components"
)

type DeletePrimaryEmailResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeletePrimaryEmailResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
