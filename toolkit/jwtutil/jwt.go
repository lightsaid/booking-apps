package jwtutil

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

const minSecretKeySize = 16

// ParseToken 方法返回的错误类型
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Maker struct {
	secretKey string
}

// NewMaker 创建一个 jwt token 管理者； secretKey 长度必须大于15
func NewMaker(secretKey string) (*Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("secretKey size invalid: 必须大于等于%d", minSecretKeySize)
	}
	return &Maker{secretKey}, nil
}

// GenToken 创建 jwt token
func (maker *Maker) GenToken(payload *JWTPayload) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return claims.SignedString([]byte(maker.secretKey))
}

// ParseToken 解释/验证 jwt token
func (maker *Maker) ParseToken(token string) (*JWTPayload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &JWTPayload{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	})
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		//  errors.Is(verr.Inner, jwt.ErrTokenExpired 无效，查看源码得：
		// fmt.Println(">>>", ok, verr.Is(jwt.ErrTokenExpired))
		if ok && verr.Is(jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	if claims, ok := jwtToken.Claims.(*JWTPayload); ok && jwtToken.Valid {
		return claims, nil
	} else {
		return nil, ErrInvalidToken
	}
}
