// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"shop/internal/dao/internal"
)

// internalRotationInfoDao is an internal type for wrapping the internal DAO implementation.
type internalRotationInfoDao = *internal.RotationInfoDao

// rotationInfoDao is the data access object for the table rotation_info.
// You can define custom methods on it to extend its functionality as needed.
type rotationInfoDao struct {
	internalRotationInfoDao
}

var (
	// RotationInfo is a globally accessible object for table rotation_info operations.
	RotationInfo = rotationInfoDao{
		internal.NewRotationInfoDao(),
	}
)

// Add your custom methods and functionality below.
