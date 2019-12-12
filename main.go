package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	fmt.Println("server start...")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/token", tokenHandler)
	http.ListenAndServe(":8080", nil)
}

// helloHandler return "Hello world!"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

// userHandler return login user's info(id & mail)
func userHandler(w http.ResponseWriter, r *http.Request) {
	mail := r.Header.Get("X-Goog-Authenticated-User-Email")
	id := r.Header.Get("X-Goog-Authenticated-User-ID")
	fmt.Fprint(w, "ID: "+id+"\nmail: "+mail)
}

// iapClaim is struct of claim.
type iapClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// authHandler
func authHandler(w http.ResponseWriter, r *http.Request) {
	// get token
	tokenStr := r.Header.Get("X-Goog-IAP-JWT-Assertion")
	audience := os.Getenv("AUDIENCE")
	claim := &iapClaim{}

	_, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		// get public key
		if _, ok := token.Header["kid"].(string); !ok {
			return nil, errors.New("not found kid")
		}
		rawKey, err := fetchPublicKey(token.Header["kid"].(string))
		if err != nil {
			return nil, err
		}
		key, err := jwt.ParseECPublicKeyFromPEM(rawKey)
		if err != nil {
			return nil, err
		}
		return key, nil
	})
	if err != nil {
		fmt.Fprint(w, "Validation: NG"+"\nerror: "+err.Error())
		return
	}

	if claim.Audience == audience {
		fmt.Fprint(w, "Validation: OK"+"\nID: "+claim.Subject+"\nmail: "+claim.Email)
		return
	}
	fmt.Fprint(w, "Validation: NG"+"\nerror: invalid audience")
}

// fetchPublicKey
func fetchPublicKey(keyID string) ([]byte, error) {
	// get publickey
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get("https://www.gstatic.com/iap/verify/public_key")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	// format publickey
	var skeys map[string]string
	if err := json.NewDecoder(r.Body).Decode(&skeys); err != nil {
		return nil, err
	}
	return []byte(skeys[keyID]), nil
}

// tokenHandler
func tokenHandler(w http.ResponseWriter, r *http.Request) {
	// get token
	tokenStr := r.Header.Get("X-Goog-IAP-JWT-Assertion")
	audience := os.Getenv("AUDIENCE")
	claim := &iapClaim{}

	_, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		// get public key
		if _, ok := token.Header["kid"].(string); !ok {
			return nil, errors.New("not found kid")
		}
		rawKey, err := fetchPublicKey(token.Header["kid"].(string))
		if err != nil {
			return nil, err
		}
		key, err := jwt.ParseECPublicKeyFromPEM(rawKey)
		if err != nil {
			return nil, err
		}
		return key, nil
	})
	if err != nil {
		fmt.Fprint(w, "Validation: NG"+"\nerror: "+err.Error())
		return
	}

	if claim.Audience == audience {
		fmt.Fprint(w, "Validation: OK"+"\nID: "+claim.Subject+"\nmail: "+claim.Email+"\ntoken: "+tokenStr)
		return
	}
	fmt.Fprint(w, "Validation: NG"+"\nerror: invalid audience")
}

