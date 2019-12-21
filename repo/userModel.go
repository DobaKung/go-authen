package repo

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	UserId string
	jwt.StandardClaims
}

type User struct {
	ID        string `gorm:"primary_key"`
	FullName  string `gorm:"not null" json:"full_name"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Token     string `sql:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// For API response
type UserPayload struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

// Constants
var ErrInvalidLoginCred = errors.New("invalid login credentials")
var ErrCantValidateLoginCred = errors.New("cannot validate login credentials")
var ErrGeneratingToken = errors.New("cannot generate login token")
