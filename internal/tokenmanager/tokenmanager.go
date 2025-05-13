package tokenmanager

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type LogtoM2MConfig struct {
	ServerURL string
	AppID     string
	AppSecret string
	Resource  string
	Scope     string
}

type LogtoTokenManager struct {
	config  *LogtoM2MConfig
	token   string
	expires time.Time
	mu      sync.Mutex
}

func NewLogtoTokenManagerDefault(config *LogtoM2MConfig) *LogtoTokenManager {
	return &LogtoTokenManager{
		config:  config,
		token:   "",
		expires: time.Now().Add(-time.Second),
	}
}

func (t *LogtoTokenManager) GetToken() (string, error) {
	if t.token != "" && t.expires.After(time.Now()) {
		return t.token, nil
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	tokenResponse, err := fetchToken(t.config)
	if err != nil {
		return "", err
	}
	t.token = tokenResponse.AccessToken
	t.expires = time.Now().Add(time.Duration(tokenResponse.ExpiresIn)*time.Second - 20*time.Second)
	return t.token, nil
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func fetchToken(config *LogtoM2MConfig) (*tokenResponse, error) {
	auth := base64.StdEncoding.EncodeToString([]byte(config.AppID + ":" + config.AppSecret))

	requestBody := url.Values{}
	requestBody.Add("grant_type", "client_credentials")
	requestBody.Add("resource", config.Resource)
	requestBody.Add("scope", config.Scope)

	request, err := http.NewRequest("POST", config.ServerURL+"/oidc/token", strings.NewReader(requestBody.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Basic "+auth)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var token tokenResponse
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
