package token

import (
	"encoding/json"
	"errors"
	"gin/constant"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

const (
	alg = "HS256"
)

type dataType map[string]interface{}

type token struct {
	expire int64 // 生成后，多少秒内有效
	before int64 // 生成前，多少秒有效（防止时间不同步）
	data dataType // 私有参数
	secret string // 生成密钥
	token string // 生成结果
}

// 自定义token负载体
type claims struct {
	jwt.StandardClaims
	Data dataType `json:"data"`
}

// 构建token结构
func New() *token {
	return &token{
		expire: 604800,
		before: 3600,
	}
}

// 设置在当前时间之后多少秒有效
func (token *token) SetExpire(expire int64) {
	token.expire = expire
}

// 设置在当前时间多少秒之前无效
func (token *token) SetBefore(before int64) {
	token.expire = before
}

// 设置生成密钥
func (token *token) SetSecret(secret string) {
	token.secret = secret
}

// 设置token私有数据
func (token *token) SetData(data dataType) {
	token.data = data
}

// 设置token字符串
func (token *token) SetToken(tokenString string) {
	token.token = tokenString
}

// 生成token
func (token *token) Encode() (string, error) {
	newToken := jwt.NewWithClaims(jwt.GetSigningMethod(alg), claims{
		jwt.StandardClaims{
			Audience:  "gin-user",                       // 受众
			ExpiresAt: time.Now().Unix() + token.expire, // 过期时间
			Id:        "gin",                            // 编号
			IssuedAt:  time.Now().Unix(),                // 签发时间
			Issuer:    "gin",                            // 签发人
			NotBefore: time.Now().Unix() - token.before, // 生效时间
			Subject:   "gin",                            // 项目
		},
		token.data,
	})

	tokenString, err := newToken.SignedString([]byte(token.secret))
	if err != nil {
		err := errors.New("生成token失败")
		return "", err
	}

	token.token = tokenString
	return tokenString, nil
}

// 解析token
func (token *token) Decode() (data dataType, state int) {
	if len(token.token) == 0 {
		state = constant.TokenNotFound
		return
	}

	newToken, err := jwt.ParseWithClaims(token.token, &claims{}, func(newToken *jwt.Token) (i interface{}, e error) {
		return []byte(token.secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				state = constant.TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				state = constant.TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				state = constant.TokenNotValidYet
			} else {
				state = constant.TokenNotValid
			}
		}
		return
	}

	if newToken != nil {
		if claims, ok := newToken.Claims.(*claims); ok && newToken.Valid {
			data = claims.Data
			token.data = data
			state = constant.TokenValid
			return
		}
	}
	state = constant.TokenNotValid
	return
}

// 纯解析token，不做任何校验
func (token *token) DecodeSegment() (data dataType, state int) {
	if len(token.token) == 0 {
		state = constant.TokenNotFound
		return
	}
	tokenSplit := strings.Split(token.token, ".")
	if len(tokenSplit) != 3 {
		state = constant.TokenNotFound
		return
	}
	tokenData, err := jwt.DecodeSegment(tokenSplit[1])
	if len(tokenData) == 0 || err != nil {
		state = constant.TokenNotValid
		return
	}

	tokenClaims := claims{}
	if err = json.Unmarshal(tokenData, &tokenClaims); err != nil {
		state = constant.TokenNotValid
		return
	}

	data = tokenClaims.Data
	token.data = data
	state = constant.TokenValid
	return
}