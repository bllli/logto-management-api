// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/bllli/logto-management-api/models/components"
)

type GetStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
