package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	UUID     uuid.UUID
	ID       uint64
	Phone    string
	NickName string
	jwt.StandardClaims
}
