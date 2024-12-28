package internal

import "github.com/uvuv-643/Web_Construct/common/proto/pkg/sso"

func ValidateAIProxyPermissions(permissions []sso.PermissionType) bool {
	for _, permission := range permissions {
		if permission.String() == "PT_SHARE" {
			return true
		}
	}
	return false
}
