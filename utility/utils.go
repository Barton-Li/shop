package utility

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"math/rand"
	"time"
)

// EncryptPassword 函数用于对密码进行加密
func EncryptPassword(password, salt string) string {
	// 首先对密码进行一次 MD5 加密
	passwordHash := gmd5.MustEncryptString(password)
	// 然后将盐值进行 MD5 加密
	saltHash := gmd5.MustEncryptString(salt)
	// 将加密后的密码和盐值拼接在一起
	combined := passwordHash + saltHash
	// 最后对拼接后的字符串进行 MD5 加密，得到最终的加密密码
	return gmd5.MustEncryptString(combined)
}

func RandInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func GetRecent7Date() (dates []string) {
	gt := gtime.New(time.Now())
	dates = []string{
		gt.Format("Y-m-d"),
		gt.Add(-gtime.D * 1).Format("Y-m-d"),
		gt.Add(-gtime.D * 2).Format("Y-m-d"),
		gt.Add(-gtime.D * 3).Format("Y-m-d"),
		gt.Add(-gtime.D * 4).Format("Y-m-d"),
		gt.Add(-gtime.D * 5).Format("Y-m-d"),
		gt.Add(-gtime.D * 6).Format("Y-m-d"),
	}
	return
}
func GetBefore7Date() (dates string) {
	gt := gtime.New(time.Now())
	dates = gt.Add(-gtime.D * 6).Format("Y-m-d")
	return
}
func GetOrderNum() (number string) {
	rand.Seed(time.Now().UnixNano())
	number = gconv.String(time.Now().UnixNano()) + gconv.String(rand.Intn(1000))
	return
}
