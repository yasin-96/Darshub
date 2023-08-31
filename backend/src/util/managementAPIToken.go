package util

import (
	"net/http"
	"strings"
)

type Response struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"bearer"`
}

func GetManagementAPIToken() (string, error) {

	url := "https://dev-l726rl1d8x1rw7du.eu.auth0.com/oauth/token"

	payload := strings.NewReader("grant_type=client_credentials&client_id=DhGgj9Y2l1GZj14p9oMCWyQv2cAqzyEF&client_secret=AtKOtFJ1f95_AQYtZmCqMX8IQnsGyEj8M8mEI5ean_mKxP676ykxeUNqyRIVsbqY&audience=https%3A%2F%2Fdev-l726rl1d8x1rw7du.eu.auth0.com%2Fapi%2Fv2%2F")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	resp := &Response{}
	err := FromJSON(resp, res.Body)
	if err != nil {
		println(err.Error())
		return "", nil
	}
	return resp.AccessToken, nil
}
