package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/util"
	"time"
)

type Claims struct {
	AppKey             string `json:"app_key"`
	AppSecret          string `json:"app_secret"`
	jwt.StandardClaims        //jwt-go 库中预定义的结构体，也是JWT的规范
}

//获取该项目的JWT Secret ，目前我们是直接使用配置文件所写入的Secret，
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

//生成token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire) //计算割刀过期时间

	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	//根据Claims结构体创建token实例，
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //备注第一次写的时候HS256写成了ES256,啊哈哈哈
	//SignedString（）：生成签名字符串，根据Secret不同，进行签名并返回标准的Token.
	//fmt.Println("GetJWTSecret:", string(GetJWTSecret()))
	token, err := tokenClaims.SignedString(GetJWTSecret()) //使用 jwt 包中的 token.SignedString() 方法对 token 进行签名并返回,传入密钥
	// 打印 token 字符串
	//fmt.Println("Generated JWT:", token)

	//fmt.Println("err:", err)
	//生成最终token
	return token, err
}
func GenerateToken2(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

//解析客服端传来的token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	//
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err

}
