package logicore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

// SignIn - Get a new token for user
func (c *Client) signIn() (*AuthResponse, error) {
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("define username and password")
	}

	data := url.Values{}
	data.Set("grant_type", c.Auth.GrantType)
	data.Set("username", c.Auth.Username)
	data.Set("password", c.Auth.Password)
	data.Set("client_id", c.Auth.ClientID)

	res, err := c.HTTPClient.PostForm(fmt.Sprintf("%s/AuthorizationServer/Access/Login", c.HostURL), data)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	if err = json.Unmarshal(body, &ar); err != nil {
		return nil, err
	}

	return &ar, nil
}
