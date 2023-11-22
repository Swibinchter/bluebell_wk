package jwt

// 用于token的生成和验证
/*
JWT是将用户信息和有效期等信息通过签名加密成一个字符串(Header.Payload.Signature)
---->Header头部：公开，描述了签名加密算法(默认是HS256)、Token类型(JWT)
---->Payload负载：公开，存放业务需要传递的数据声明(claim)，默认字段包括签发人、过期时间、主题、受众、生效时间、签发时间、编号，也可以自定义字段
---->Signature签名：加密，将前面两部分和salt一起进行签名加密后生成的字符串
*/

import (
	"errors"
	"fmt"
	"goWebCli/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

var ErrInvalidToken = errors.New("invalid token")

// 在默认官方字段之外增加一个username
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成Token
func GenToken(userID int64, username string) (string, error) {
	// 创建token中的声明
	c := MyClaims{
		userID,
		username,
		jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(time.Duration(setting.Config.JWTExpire * int64(time.Hour))).Unix(),
			// 签发人
			Issuer: "bluebell",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用salt签名获得完整的Token字符串
	return token.SignedString([]byte("swibinchter"))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token，成功的话将token中的声明存放到结构体变量mc中
	mc := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(t *jwt.Token) (interface{}, error) {
		return []byte(setting.Config.Salt), nil
	})
	fmt.Printf("获取到的token是%v\n", token)
	// token.Valid表示token是否通过验证
	if token == nil || !token.Valid {
		zap.L().Error("invalid token", zap.Error(err))
		return nil, ErrInvalidToken
	} else if err != nil {
		zap.L().Error("parse token failed", zap.Error(err))
		return nil, err
	}
	// token通过验证且err==nil
	return mc, nil
}
