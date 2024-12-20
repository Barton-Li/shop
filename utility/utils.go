package utility

import "github.com/gogf/gf/v2/crypto/gmd5"

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
