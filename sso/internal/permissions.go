package internal

import "github.com/uvuv-643/Web_Construct/common/proto/pkg/sso"

func GetPermissionType(permission string) sso.PermissionType {
	if permission == "PT_READ" {
		return sso.PermissionType_PT_READ
	} else if permission == "PT_WRITE" {
		return sso.PermissionType_PT_WRITE
	} else if permission == "PT_DELETE" {
		return sso.PermissionType_PT_DELETE
	} else if permission == "PT_SHARE" {
		return sso.PermissionType_PT_SHARE
	} else if permission == "PT_AUDIT" {
		return sso.PermissionType_PT_AUDIT
	} else if permission == "PT_MANAGE" {
		return sso.PermissionType_PT_MANAGE
	} else {
		return sso.PermissionType_PT_UNDEFINED
	}
}
