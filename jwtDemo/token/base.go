package jwt

import "time"

type TokenGenerator interface {
	GenerateToken(code string, expire time.Duration) (string, error)
}
