package googlessostore

import (
	"context"
	"encoding/json"
	"fmt"
	"foodlive/modules/authsso/googlesso/googlessomodel"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"time"
)

func (*authSsoStore) ValidateGoogleJwt(ctx context.Context, tokenString string) (*googlessomodel.GoogleClaims, error) {
	claimsStruct := googlessomodel.GoogleClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
			if err != nil {
				return nil, err
			}
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*googlessomodel.GoogleClaims)
	if !ok {
		return nil, googlessomodel.ErrInvalidGoogleJwt
	}

	if claims.Issuer != "accounts.google.com" {
		if claims.Issuer != "https://accounts.google.com" {
			return nil, googlessomodel.ErrInvalidJwtIss
		}
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, googlessomodel.ErrGoogleJwtExpire
	}

	return claims, nil
}

func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}
	key, ok := myResp[keyID]
	if !ok {
		return "", googlessomodel.ErrGoogleKeyNotFound
	}
	return key, nil
}
