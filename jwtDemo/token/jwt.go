package jwt

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

func NewJWTTokenGen(Issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		privateKey: privateKey,
		issuer:     Issuer,
		nowFunc:    time.Now,
	}
}

func (J JWTTokenGen) GenerateToken(code string, expire time.Duration) (string, error) {
	nowTime := J.nowFunc()
	tkn := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.RegisteredClaims{
		Issuer:    J.issuer,
		IssuedAt:  jwt.NewNumericDate(nowTime),
		ExpiresAt: jwt.NewNumericDate(nowTime.Add(expire)),
		Subject:   code,
	})
	return tkn.SignedString(J.privateKey)
}
