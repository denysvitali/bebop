package oauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type provider struct {
	config  *oauth2.Config
	getUser func(*http.Client) (*user, error)
}

type providerConfig struct {
	endpoint oauth2.Endpoint
	scopes   []string
	getUser  func(*http.Client) (*user, error)
}

type user struct {
	id   string
	name string
}

var providerConfigs = map[string]providerConfig{
	"google": {
		endpoint: google.Endpoint,
		scopes:   []string{"profile"},
		getUser:  getOauthUser("https://www.googleapis.com/oauth2/v2/userinfo"),
	},
	"facebook": {
		endpoint: facebook.Endpoint,
		scopes:   []string{"public_profile"},
		getUser:  getOauthUser("https://graph.facebook.com/me?fields=id,name"),
	},
	"github": {
		endpoint: github.Endpoint,
		scopes:   []string{},
		getUser:  getOauthUser("https://api.github.com/user"),
	},
}

func getOauthUser(userEndpoint string) func(c *http.Client) (*user, error) {
	return func(c *http.Client) (*user, error){
		u := struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		}{}

		err := getJSON(c, userEndpoint, &u)
		if err != nil {
			return nil, err
		}

		return &user{id: strconv.FormatInt(u.ID, 10), name: u.Name}, nil
	}
}

func getJSON(c *http.Client, url string, v interface{}) error {
	response, err := c.Get(url)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return fmt.Errorf("bad request status code: %v", response.StatusCode)
	}

	err = json.NewDecoder(io.LimitReader(response.Body, 1<<20)).Decode(v)
	if err != nil {
		return fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return nil
}
