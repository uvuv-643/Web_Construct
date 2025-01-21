package internal

import (
	"github.com/uvuv-643/Web_Construct/common/proto/pkg/sso"
)

func BuildUserPermissions(user *User) []*sso.AppPermission {
	userPermissions := make(map[string][]sso.PermissionType)
	for _, role := range user.Roles {
		userPermissions[role.ApplicationID] = append(
			userPermissions[role.ApplicationID],
			GetPermissionsType(string(role.Role)),
		)
	}
	appPermissions := make([]*sso.AppPermission, 0)
	for key, permissions := range userPermissions {
		appPermissions = append(appPermissions, &sso.AppPermission{
			AppUuid:     key,
			Permissions: permissions,
		})
	}
	return appPermissions
}
