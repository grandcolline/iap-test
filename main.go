package main

import (
	"fmt"
	// "time"
	"net/http"
	"strconv"
	// "os"
	// "encoding/json"

	// jwt "github.com/dgrijalva/jwt-go"
	// "github.com/imkira/gcp-iap-auth/jwt"
)

const(
	PublicKeysURL = "https://www.gstatic.com/iap/verify/public_key"
)

func main() {
	fmt.Println("server start...")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/auth", authHandler)
	http.ListenAndServe(":8080", nil)
}

// helloHandler return "Hello world!"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	mail := r.Header.Get("X-Goog-Authenticated-User-Email")
	id := r.Header.Get("X-Goog-Authenticated-User-ID")
	fmt.Fprint(w, "ID: "+id+"\nmail: "+mail)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	auth := false
	token := r.Header.Get("X-Goog-IAP-JWT-Assertion")
	// audience := os.Getenv("audience")

	// aud, _ := jwt.ParseAudience(audience)
	// publicKeys, _ := jwt.FetchPublicKeys()
	// cfg := &jwt.Config{
	// 	Audiences:  []*jwt.Audience{aud},
	// 	PublicKeys: publicKeys,
	// }

	// if err := jwt.ValidateRequestClaims(r, cfg); err != nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// } else {
	// 	w.WriteHeader(http.StatusOK)
	// }

	fmt.Fprint(w, "auth: "+strconv.FormatBool(auth)+"\ntoken:" + token)
}

// type Claims struct {
// 	jwt.StandardClaims
// 	Email string `json:"email,omitempty"`
// 
// 	cfg *Config
// }
// 
// type Config struct {
// 	PublicKeys     map[string]PublicKey
// 	MatchAudiences *regexp.Regexp
// }
// 
// // fetchPublicKeys 公開鍵の取得
// func fetchPublicKeys() (map[string][]byte, error) {
// 	// get publickey
// 	client := &http.Client{Timeout: 10 * time.Second}
// 	r, err := client.Get(PublicKeysURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer r.Body.Close()
// 
// 	// format publickey
// 	var skeys map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&skeys); err != nil {
// 		return nil, err
// 	}
// 	bkeys := make(map[string][]byte)
// 	for k, v := range skeys {
// 		if len(v) != 0 {
// 			bkeys[k] = []byte([]byte(v))
// 		}
// 	}
// 	return bkeys, nil
// }
