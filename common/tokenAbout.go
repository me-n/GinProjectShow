package common

import (
	"GinProjectShow/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)
//token相关函数

var hmacKey = []byte("me-n_secret_key") //HMAC签名密钥，非常重要！

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//发放token
func ReleaseToken(user model.User) (string, error) {
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(), //设置token过期时间为一周
			IssuedAt:  time.Now().Unix(),                         //设置发布时间为当前时间
			Issuer:    "me-n",                                    //设置发布者为MN
			Subject:   "HMAC_UserToken",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //设置加密方法为HMAC256并生成token
	signedString, err := token.SignedString(hmacKey)           //利用之前设置的key对token进行签名
	if err != nil {
		return "", err
	}
	return signedString, nil
}

//进行token验证
func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return hmacKey, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, claims, nil
}
