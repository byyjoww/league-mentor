package auth

import (
	"fmt"
	"net/http"
)

type ApiKey struct {
	key string
}

func NewApiKey(apiKey string) *ApiKey {
	return &ApiKey{
		key: apiKey,
	}
}

func (a *ApiKey) SetAuth(req *http.Request) {
	req.Header.Set("X-Riot-Token", a.key)
}

func (a *ApiKey) String() string {
	return fmt.Sprintf("%s:%s", "X-Riot-Token", a.key)
}
