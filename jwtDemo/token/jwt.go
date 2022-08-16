package token

import (
	"crypto/rsa"
	"time"
)
import "github.com/golang-jwt/jwt/v4"

type JWTTokenGen struct {
	privateKey *rsa.PrivateKey
	issuer     string
	nowFunc    func() time.Time
}

func NewJWTTokenGen(Issuer string, privateKey *rsa.PrivateKey, nowFunc func() time.Time) *JWTTokenGen {
	return &JWTTokenGen{
		privateKey: privateKey,
		issuer:     Issuer,
		nowFunc:    nowFunc,
	}
}

func (t *JWTTokenGen) GenerateToken(code string, expire time.Duration) (string, error) {
	nowTime := t.nowFunc()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.RegisteredClaims{
		Issuer:    t.issuer,
		IssuedAt:  jwt.NewNumericDate(nowTime),
		ExpiresAt: jwt.NewNumericDate(nowTime.Add(expire)),
		Subject:   code,
	})
	return tkn.SignedString(t.privateKey)
}
