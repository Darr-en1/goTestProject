package token

import "time"

type Generator interface {
	GenerateToken(code string, expire time.Duration) (string, error)
}
