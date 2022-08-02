package logicore

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Base endpoint of the url
type Client struct {
	HostURL      string
	HTTPClient   *http.Client
	AccessToken  string
	RefreshToken string
	Auth         AuthStruct
}

type AuthStruct struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	ClientID  string `json:"client_id"`
	GrantType string `json:"grant_type"`
}

type AuthResponse struct {
	ActingOwnerID int    `json:"actionOwnerId"`
	RefreshToken  string `json:"refresh_token"`
	AccessToken   string `json:"access_token"`
}

func NewClient(
	host string,
	username string,
	password string,
	clientID string,
) (*Client, error) {
	client := Client{
		HTTPClient: &http.Client{Timeout: 100 * time.Second},
		HostURL:    host,
	}

	// Check if username or password are provided
	if username == "" || password == "" {
		return &Client{}, nil
	}

	client.Auth = AuthStruct{
		GrantType: "password",
		Username:  username,
		Password:  password,
		ClientID:  clientID,
	}

	ar, err := client.signIn()
	if err != nil {
		return nil, err
	}

	client.AccessToken = ar.AccessToken
	client.RefreshToken = ar.RefreshToken

	return &client, nil
}

func (c *Client) doRequest(
	req *http.Request,
	authToken *string,
) ([]byte, error) {
	token := c.AccessToken

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
