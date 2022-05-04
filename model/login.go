package model

import "github.com/dgrijalva/jwt-go"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
