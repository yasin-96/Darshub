package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/auth0/go-auth0/management"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	Expiring    int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func GetManagementClient() (management.Management, error) {
	domain := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	auth0API, err := management.New(
		domain,
		management.WithClientCredentials(context.Background(), clientID, clientSecret),
	)
	if err != nil {
		return management.Management{}, err
	}

	return *auth0API, nil
}

func GetAccesToken() (string, error) {
	domain := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	url := fmt.Sprintf("https://%s/oauth/token", domain)

	payloadString := "grant_type=client_credentials&client_id=" + clientID + "&client_secret=" + clientSecret + "&audience=https%3A%2F%2F" + domain + "%2Fapi%2Fv2%2F"

	payload := strings.NewReader(payloadString)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	//defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	tokenResponse := TokenResponse{}
	err := json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}
	return tokenResponse.AccessToken, nil
}
