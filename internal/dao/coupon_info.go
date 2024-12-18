// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"shop/internal/dao/internal"
)

// internalCouponInfoDao is an internal type for wrapping the internal DAO implementation.
type internalCouponInfoDao = *internal.CouponInfoDao

// couponInfoDao is the data access object for the table coupon_info.
// You can define custom methods on it to extend its functionality as needed.
type couponInfoDao struct {
	internalCouponInfoDao
}

var (
	// CouponInfo is a globally accessible object for table coupon_info operations.
	CouponInfo = couponInfoDao{
		internal.NewCouponInfoDao(),
	}
)

// Add your custom methods and functionality below.
