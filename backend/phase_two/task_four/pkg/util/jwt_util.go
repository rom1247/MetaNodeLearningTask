package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims 自定义 payload
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT Token
func GenerateToken(userID uint, secretKey string, expireMinutes time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireMinutes * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "myapp",
			Subject:   "user_token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ParseToken 解析 Token 并验证
func ParseToken(tokenStr string, secretKey string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
