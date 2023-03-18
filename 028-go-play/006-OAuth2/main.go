package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"

)

type Client struct {
	config *oauth2.Config
	token  *oauth2.Token
}

func main() {
	// Set up the OAuth2 config
	config := &oauth2.Config{
		ClientID:     "your_client_id",
		ClientSecret: "your_client_secret",
		Scopes:       []string{"read", "write"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://example.com/oauth2/auth",
			TokenURL: "https://example.com/oauth2/token",
		},
	}

	// constructing the client struct
	client := &Client{
		config: config,
	}

	// Get the authorization URL and prompt the user to visit it
	authURL, err := client.GetAuthURL()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Visit the URL for the auth dialog: %v\n", authURL)

	// Wait for the user to authorize the app and get the authorization code
	var code string
	fmt.Print("Enter the authorization code: ")
	fmt.Scan(&code)

	// Exchange the authorization code for an access token
	err = client.ExchangeAuthCode(context.Background(), code)
	if err != nil {
		log.Fatal(err)
	}

	// Use the access token to make an authenticated request
	resp, err := client.MakeRequest("https://example.com/api/v1/resource")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Print the response body
	fmt.Println(resp.Status)
}

// GetAuthURL gets the authorization URL and returns it to the caller
func (c *Client) GetAuthURL() (string, error) {
	authURL := c.config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return authURL, nil
}

// ExchangeAuthCode exchanges the authorization code for an access token
func (c *Client) ExchangeAuthCode(ctx context.Context, code string) error {
	token, err := c.config.Exchange(ctx, code)
	if err != nil {
		return err
	}
	c.token = token
	return nil
}

// MakeRequest makes an authenticated HTTP request
func (c *Client) MakeRequest(url string) (*http.Response, error) {
	client := c.config.Client(context.Background(), c.token)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
