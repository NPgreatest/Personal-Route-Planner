package utils

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claim struct {
	UserId int
	Name   string
	jwt.StandardClaims
}

var jwtKey = []byte("gpa-manager-npgreatest")

func CreateToken(id int, name string, expireDuration time.Duration) (string, error) {
	claims := &Claim{
		UserId: id,
		Name:   name,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "mgh",
			Subject:   "User_Token",
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func VerifyToken(token string) (int, string, bool) {
	if token == "" {
		return 0, "", false
	}

	tok, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Println("ParseWithClaims error %v", err)
		return 0, "", false
	}

	if claims, ok := tok.Claims.(*Claim); ok && tok.Valid {
		return claims.UserId, claims.Name, true
	} else {
		fmt.Println("%v", err)
		return 0, "", false
	}

}

func GenerateId(x int64) int64 {
	node, err := snowflake.NewNode(x)
	if err != nil {
		return 0
	}
	return node.Generate().Int64()
}
