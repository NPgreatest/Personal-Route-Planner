package utils

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"strings"
	"time"
)

type Claim struct {
	Name string
	jwt.StandardClaims
}

var jwtKey = []byte("gpa-manager-npgreatest")

func CreateToken(name string, expireDuration time.Duration) (string, error) {
	claims := &Claim{
		Name: name,
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

func VerifyToken(token string) (string, bool) {
	if token == "" {
		return "", false
	}

	tok, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Println("ParseWithClaims error %v", err)
		return "", false
	}

	if claims, ok := tok.Claims.(*Claim); ok && tok.Valid {
		return claims.Name, true
	} else {
		fmt.Println("%v", err)
		return "", false
	}

}

func GenerateId(x int64) int64 {
	node, err := snowflake.NewNode(x)
	if err != nil {
		return 0
	}
	return node.Generate().Int64()
}

func WriteImages(data string) (string, error) {
	DataArr := strings.Split(data, ",")
	imgBase64 := DataArr[1]
	imgs, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return "", err
	}
	timenow := time.Now().Unix()
	imgname := strconv.FormatInt(timenow, 10) + ".jpg"
	file, err := os.OpenFile("./images/"+imgname, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	w := bufio.NewWriter(file)
	_, err = w.WriteString(string(imgs))
	if err != nil {
		return "", err
	}
	w.Flush()
	defer file.Close()
	return "http://localhost:8080/images/" + imgname, nil
}
