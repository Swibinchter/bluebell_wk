package md5

/*
使用MD5哈希算法对明文密码进行加密
MD5可以将任意长度的数据转换成固定长度128位的哈希值，具有固定长度、不可逆性、雪崩效应(输入数据的微小变化会导致输出大变)
需要注意的是，MD5存在碰撞攻击等安全性问题，如有需求可以使用更安全的SHA-256
*/

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5Encode 对字符串进行MD5加密，返回哈希值字符串
func Md5Encode(data string) string {
	// 创建一个新的MD5哈希值实例
	h := md5.New()
	// 使用h.Write()方法添加字符串，可以多次添加，但字符串需要转换成字节切片
	h.Write([]byte(data))
	// h.Sum()方法计算哈希值，以字节切片形式保存到参数中，如果参数为nil，则直接返回
	cipherStr := h.Sum(nil)
	// 将哈希值的格式从字节切片转换成十六进制字符串（小写）
	return hex.EncodeToString(cipherStr)
	// 将哈希值的格式从字节切片转换成十六进制字符串（大写）
	// return strings.ToUpper(hex.EncodeToString(cipherStr))
}

// EncryptPassword 对明文密码添加salt后进行MD5加密
func EncryptPassword(plainPassword, salt string) string {
	return Md5Encode(plainPassword + salt)
}

// ValidatePassWord 验证密码是否正确
func ValidatePassWord(plainPassword, salt, password string) bool {
	return Md5Encode(plainPassword+salt) == password
}
