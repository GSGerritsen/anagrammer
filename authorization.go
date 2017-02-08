package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	privateKeyPath = "./keys/app.rsa"
	publicKeyPath  = "./keys/app.rsa.pub"
)

type ClaimsSet struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var VerifyKey, SignKey []byte

func initKeys() {
	var err error

	SignKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	VerifyKey, err = ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal(err)
		return
	}
}

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

func TestProtectedAccess(w http.ResponseWriter, r *http.Request) error {
	response := Response{"Gained access to protected resource"}
	JsonResponse(response, w)
	return nil
}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var u User
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	email := u.Email
	password := u.Password

	// Check database here
	if strings.ToLower(email) != "testEmail@gmail.com" {
		if password != "pw" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return nil
		}
	}

	claims := ClaimsSet{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    "testIssuer",
			Subject:   "Hi",
		},
	}

	// Used to be: claims := make(jwt.MapClaims) => claims["iss"] = "testClaim"
	signer := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)

	parsedPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(SignKey)
	if err != nil {
		panic(err)
	}
	tokenString, err := signer.SignedString(parsedPrivateKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
	}

	jwtCookie := http.Cookie{Name: "jwt_token", Value: tokenString, Expires: time.Now().Add(time.Minute * 10)}
	http.SetCookie(w, &jwtCookie)

	http.Redirect(w, r, "/vue-tests", 200)
	return nil

}

func HandleLogout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	deleteCookie := http.Cookie{Name: "jwt_token", Value: "none", Expires: time.Now()}
	http.SetCookie(w, &deleteCookie)
	http.Redirect(w, r, "/vue-tests", 200)
}

func HandleTokenAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := ExtractAndValidate(w, r)
	if err != nil {
		fmt.Errorf("Something bad happened: %s", err)
	}
	http.Redirect(w, r, "/vue-tests", 200)
}

func ExtractAndValidate(w http.ResponseWriter, r *http.Request) error {

	result, err := ExtractToken(r)
	if err != nil {
		http.Redirect(w, r, "/vue-tests", 401)
	}
	ValidateToken(result)
	fmt.Println("Success in reading token")
	return nil
}

func ExtractToken(r *http.Request) (string, error) {
	/*
		tokenHeader := strings.Split(r.Header.Get("Authorization"), " ")
		if tokenHeader[0] == "Bearer" && len(tokenHeader) == 2 {
			return tokenHeader[1], nil
		}
	*/

	jwtCookie, err := r.Cookie("jwt_token")
	if err != nil {
		return "", fmt.Errorf("no token found")
	}
	return jwtCookie.Value, nil

}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	if len(tokenString) == 0 {
		return nil, fmt.Errorf("No token received")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return SignKey, nil
	})

	if err != nil {
		return token, nil
	}
	fmt.Println("VERIFIED TOKEN")
	return nil, fmt.Errorf("RSA private or RSA public not set, check those")
}
