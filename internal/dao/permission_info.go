// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"shop/internal/dao/internal"
)

// internalPermissionInfoDao is an internal type for wrapping the internal DAO implementation.
type internalPermissionInfoDao = *internal.PermissionInfoDao

// permissionInfoDao is the data access object for the table permission_info.
// You can define custom methods on it to extend its functionality as needed.
type permissionInfoDao struct {
	internalPermissionInfoDao
}

var (
	// PermissionInfo is a globally accessible object for table permission_info operations.
	PermissionInfo = permissionInfoDao{
		internal.NewPermissionInfoDao(),
	}
)

// Add your custom methods and functionality below.