// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"shop/internal/dao/internal"
)

// internalConsigneeInfoDao is an internal type for wrapping the internal DAO implementation.
type internalConsigneeInfoDao = *internal.ConsigneeInfoDao

// consigneeInfoDao is the data access object for the table consignee_info.
// You can define custom methods on it to extend its functionality as needed.
type consigneeInfoDao struct {
	internalConsigneeInfoDao
}

var (
	// ConsigneeInfo is a globally accessible object for table consignee_info operations.
	ConsigneeInfo = consigneeInfoDao{
		internal.NewConsigneeInfoDao(),
	}
)

// Add your custom methods and functionality below.