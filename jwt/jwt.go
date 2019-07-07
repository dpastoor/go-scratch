package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/go-homedir"
)

func main() {
	// when looking at https://github.com/probot/probot/blob/master/src/private-key.ts
	// can see that the signing key should ultimately be a string
	fp, _ := homedir.Expand("~/Downloads/kb.2019-07-07.private-key.pem")
	bytes, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
		panic("error reading pem file")
	}
	rsaKey, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		fmt.Println(err)
		panic("error parsing pem file")
	}
	enc := b64.StdEncoding.EncodeToString(bytes)
	fmt.Println(string(enc))

	// Create the Claims
	claims := &jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		Issuer:    "1",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(rsaKey)
	fmt.Println("encoded jwt")
	fmt.Printf("%v %v", ss, err)
}
