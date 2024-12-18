// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"shop/internal/dao/internal"
)

// internalRolePermissionInfoDao is an internal type for wrapping the internal DAO implementation.
type internalRolePermissionInfoDao = *internal.RolePermissionInfoDao

// rolePermissionInfoDao is the data access object for the table role_permission_info.
// You can define custom methods on it to extend its functionality as needed.
type rolePermissionInfoDao struct {
	internalRolePermissionInfoDao
}

var (
	// RolePermissionInfo is a globally accessible object for table role_permission_info operations.
	RolePermissionInfo = rolePermissionInfoDao{
		internal.NewRolePermissionInfoDao(),
	}
)

// Add your custom methods and functionality below.