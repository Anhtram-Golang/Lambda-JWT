package main
//Create and validate JSON web token

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
  "github.com/dgrijalva/jwt-go"
  "time"
	"log"
	"os"
)

//When creating and validating JSON web token, the payload should specify the name of a secret key to use in the process
//action: "create" or "verify"
//for create: secretname,action,userid are required
//for verify: secretname, action, tokenstr are required
type Payload struct {
	SecretName string `json:"secretkeyname"`
	Action     string `json:"action"`
	UserID     uint64 `json:"userid"`
	TokenStr   string `json:"tokenstr"`
}

func main() {

		if len(os.Args) >= 2 && os.Args[1] == "test" {
			err := LocalTest()

			if err != nil {
				fmt.Println(err)
			}
		} else {
			lambda.Start(handleRequest)
		}
}


func LocalTest() error {
	var testPayload = Payload{
		SecretName: "JWTsecret",
		Action    : "create",
		UserID    : 1,
	}

	log.Println("LocalTest() create starts----")
	tokenStr, err := handleRequest(testPayload)
	log.Println("LocalTest() tokenStr=", tokenStr)
	log.Println("LocalTest() err=", err)

	if err == nil {
		log.Println("LocalTest() verify starts, using the newly create jwt ----")
		testPayload.TokenStr = tokenStr
		testPayload.Action = "verify"
		log.Println("LocalTest() testPayload=", testPayload)
		result, errV := handleRequest(testPayload)
		log.Println("LocalTest() errV=", errV)
		log.Println("LocalTest() result=", result)
	}

	return nil
}



//When creating and validating jwt, the payload should specify the name of a secret key to use in the process
//for create: secretname,action,userid are required
//for verify: secretname, action, tokenstr are required
func handleRequest(payload Payload) (string, error) {
  action := payload.Action
	var result = ""
	var err error

	if action == "create" {
		result, err = CreateToken(payload.UserID, payload.SecretName)
		if err != nil {
			log.Println("Error: " + err.Error())
			return "", err
		}
	} else if action == "verify" {
		result, err = VerifyToken(payload.TokenStr, payload.SecretName)
		if err != nil {
			log.Println("Error: " + err.Error())
			return "", err
		}
	}
	return result, err
}

//creates a new json web token
func CreateToken(userId uint64, secret_name string) (string, error) {

  //Retrieve secret value from secrets manager
	secret, err := getSecretValue(secret_name);
	if err != nil {
		return "", err
	}
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userId
  atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
  if err != nil {
     return "", err
  }
	log.Println("Token is successfully created")
  return token, nil
}

//verifies if a string tokenStr is valid
func VerifyToken(tokenStr string, secret_name string) (string, error) {
	 var result = ""
	 //Retrieve secret value from secrets manager
	 secret, err := getSecretValue(secret_name);
	 verifyToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		 return[]byte(secret), nil
	 })
	 if err == nil  && verifyToken.Valid{
		 result = "Valid"
	 } else {
		 result = "Invalid"
	 }
	 log.Println("VerifyToken result =", result)

	 return result, err
}
