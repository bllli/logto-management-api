package logtomanagementapi

import (
	"context"

	"github.com/bllli/logto-management-api/internal/tokenmanager"
	"github.com/bllli/logto-management-api/models/components"
)

func NewLogtoTokenManagerWithClient(config *tokenmanager.LogtoM2MConfig) *LogtoManagementAPI {
	tokenManager := tokenmanager.NewLogtoTokenManagerDefault(config)
	return New(
		WithServerURL(config.ServerURL),
		WithSecuritySource(func(ctx context.Context) (components.Security, error) {
			token, err := tokenManager.GetToken()
			if err != nil {
				return components.Security{}, err
			}
			return components.Security{
				BearerAuth: &token,
			}, nil
		}),
	)
}
