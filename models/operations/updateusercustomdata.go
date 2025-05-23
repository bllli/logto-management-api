// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/bllli/logto-management-api/models/components"
)

// UpdateUserCustomDataCustomData - Partial custom data object to update for the given user ID.
type UpdateUserCustomDataCustomData struct {
}

type UpdateUserCustomDataRequestBody struct {
	// Partial custom data object to update for the given user ID.
	CustomData UpdateUserCustomDataCustomData `json:"customData"`
}

func (o *UpdateUserCustomDataRequestBody) GetCustomData() UpdateUserCustomDataCustomData {
	if o == nil {
		return UpdateUserCustomDataCustomData{}
	}
	return o.CustomData
}

type UpdateUserCustomDataRequest struct {
	// The unique identifier of the user.
	UserID      string                          `pathParam:"style=simple,explode=false,name=userId"`
	RequestBody UpdateUserCustomDataRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateUserCustomDataRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *UpdateUserCustomDataRequest) GetRequestBody() UpdateUserCustomDataRequestBody {
	if o == nil {
		return UpdateUserCustomDataRequestBody{}
	}
	return o.RequestBody
}

type UpdateUserCustomDataResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Updated custom data in JSON for the given user ID.
	Object map[string]any
}

func (o *UpdateUserCustomDataResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateUserCustomDataResponse) GetObject() map[string]any {
	if o == nil {
		return nil
	}
	return o.Object
}
