package util

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"bearer"`
}

func GetManagementAPIToken() (string, error) {
	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	url := fmt.Sprintf("https://%s/oauth/token", auth0Domain)
	audience := "https%3A%2F%2Fdev-l726rl1d8x1rw7du.eu.auth0.com%2Fapi%2Fv2%2F"

	formattedUrl := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&audience=%s", clientId, clientSecret, audience)

	payload := strings.NewReader(formattedUrl)

	req, reqErr := http.NewRequest("POST", url, payload)
	if reqErr != nil {
		return "", reqErr
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		return "", resErr
	}

	defer res.Body.Close()

	resp := &Response{}
	err := FromJSON(resp, res.Body)
	if err != nil {
		println(err.Error())
		return "", nil
	}
	if resp.AccessToken == "" {
		return "", errors.New("token is not valid")
	}

	return resp.AccessToken, nil
}
