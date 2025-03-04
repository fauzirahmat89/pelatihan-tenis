package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("dadwdwfsafasdsadsa")
var AdminJWT_KEY = []byte("dbasdbusabduasudb165487941165vbasdadsa121321]/]/]/]")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}