package repo

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/teris-io/shortid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
)

// ID generator
var sid, _ = shortid.New(1, shortid.DefaultABC, 2345)

func (user User) Create() error {
	// Generate user id
	var sidErr error
	user.ID, sidErr = sid.Generate()
	if sidErr != nil {
		log.Println("User.Create: " + sidErr.Error())
		return errors.New("cannot generate user id")
	}

	// Hash password
	hashedPwd, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		log.Println("User.Create: " + hashErr.Error())
		return errors.New("cannot create user")
	}
	user.Password = string(hashedPwd)

	// Write to DB
	if writeErr := GetDB().Create(&user).Error; writeErr != nil {
		log.Println("User.Create: " + writeErr.Error())
		if strings.Contains(writeErr.Error(), "1062") { // error code for duplicate entry
			return errors.New("user already exists")
		}
		return errors.New("cannot create user")
	}
	log.Println("User.Create: " + user.Email + " created")
	return nil
}

func Login(email, pwd string) (token string, err error) {
	var user User

	// Check if user exists
	if findErr := GetDB().Where("email = ?", email).First(&user).Error; findErr != nil {
		log.Println("User.Login: " + findErr.Error())
		return "", ErrInvalidLoginCred
	}

	// Compare password
	bcrErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	if bcrErr != nil {
		if bcrErr == bcrypt.ErrMismatchedHashAndPassword {
			return "", ErrInvalidLoginCred
		}
		log.Println("User.Login: " + bcrErr.Error())
		return "", ErrCantValidateLoginCred
	}

	// Generate JWT
	var jwtErr error
	tk := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), Token{UserId: user.ID})
	token, jwtErr = tk.SignedString([]byte(os.Getenv("token_secret")))
	if jwtErr != nil {
		log.Println("User.Login: " + jwtErr.Error())
		return "", ErrGeneratingToken
	}
	return token, nil
}
