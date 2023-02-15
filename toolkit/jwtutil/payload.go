package jwtutil

import (
	"strconv"
	"time"
	"toolkit/random"

	"github.com/golang-jwt/jwt/v4"
)

const defaultExpiresAt = 15

// JWTToken 负载数据
// 包含数据： uid 唯一标识，jwt.RegisteredClaims
type JWTPayload struct {
	UID int64 `json:"uid"`
	jwt.RegisteredClaims
}

// NewJWTPayload 创建一个Token Payload，
// 如果claims.ID不存在则默认生成一个
// 如果claims.ExpiresAt 不存在 默认15分钟
func NewJWTPayload(uid int64, claims jwt.RegisteredClaims) *JWTPayload {
	if claims.ID == "" {
		claims.ID = strconv.FormatInt(time.Now().UnixNano(), 10) + random.RandomString(8)
	}
	if claims.ExpiresAt == nil {
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(defaultExpiresAt * time.Minute))
	}
	if claims.IssuedAt == nil {
		claims.IssuedAt = jwt.NewNumericDate(time.Now())
	}
	return &JWTPayload{
		UID:              uid,
		RegisteredClaims: claims,
	}
}
