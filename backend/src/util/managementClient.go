package util

import (
	"context"
	"os"

	"github.com/auth0/go-auth0/management"
)

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
