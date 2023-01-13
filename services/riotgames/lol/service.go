package lol

import (
	"context"
	"fmt"
	"net/http"

	"github.com/byyjoww/league-mentor/config"
	"github.com/byyjoww/league-mentor/services/http/client"
	"github.com/byyjoww/league-mentor/services/http/client/auth"
	"github.com/byyjoww/league-mentor/services/http/decoder"
)

const (
	northAmericaBaseUrl = "https://na1.api.riotgames.com"
	americasBaseUrl     = "https://americas.api.riotgames.com"
	spectatorV4         = "/lol/spectator/v4/active-games/by-summoner/{encryptedSummonerId}"
	matchV5             = "/lol/match/v5/matches/{matchId}"
	matchesV5           = "/lol/match/v5/matches/by-puuid/{puuid}/ids"
	summonerV4          = "/lol/summoner/v4/summoners/by-name/{summonerName}"
)

const (
	userAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
	acceptLanguage = "en-US,en;q=0.9,pt-BR;q=0.8,pt;q=0.7"
	acceptCharset  = "application/x-www-form-urlencoded; charset=UTF-8"
	origin         = "https://developer.riotgames.com"
)

var (
	headers = map[string]string{
		"User-Agent":      userAgent,
		"Accept-Language": acceptLanguage,
		"Accept-Charset":  acceptCharset,
		"Origin":          origin,
	}
)

type Client interface {
	GetSummonerByName(ctx context.Context, name string) (Summoner, error)
	GetMatchesByPuuid(ctx context.Context, puuid string) ([]string, error)
	GetMatchById(ctx context.Context, matchId string) (Match, error)
}

type CustomClient struct {
	client client.Client
}

func NewCustomClient(cfg config.RiotGames) *CustomClient {
	return &CustomClient{
		client: newHttpClient(cfg.ApiKey),
	}
}

func (c *CustomClient) GetSummonerByName(ctx context.Context, name string) (Summoner, error) {
	var (
		summoner Summoner
		params   = map[string]string{
			"{summonerName}": name,
		}
	)

	statusCode, err := c.client.Get(northAmericaBaseUrl + summonerV4).
		WithContext(ctx).
		Replace(params).
		WithHeaders(headers).
		DoAndUnmarshal(&summoner)

	if err == nil && statusCode >= 300 {
		err = fmt.Errorf("request failed with status code: %d", statusCode)
	}

	return summoner, err
}

func (c *CustomClient) GetMatchesByPuuid(ctx context.Context, puuid string) ([]string, error) {
	var (
		matches []string
		params  = map[string]string{
			"{puuid}": puuid,
		}
	)

	statusCode, err := c.client.Get(americasBaseUrl + matchesV5).
		WithContext(ctx).
		Replace(params).
		WithHeaders(headers).
		DoAndUnmarshal(&matches)

	if err == nil && statusCode >= 300 {
		err = fmt.Errorf("request failed with status code: %d", statusCode)
	}

	return matches, err
}

func (c *CustomClient) GetMatchById(ctx context.Context, matchId string) (Match, error) {
	var (
		matches Match
		params  = map[string]string{
			"{matchId}": matchId,
		}
	)

	statusCode, err := c.client.Get(americasBaseUrl + matchV5).
		WithContext(ctx).
		Replace(params).
		WithHeaders(headers).
		DoAndUnmarshal(&matches)

	if err == nil && statusCode >= 300 {
		err = fmt.Errorf("request failed with status code: %d", statusCode)
	}

	return matches, err
}

func newHttpClient(apiKey string) client.Client {
	var (
		httpClient = http.DefaultClient
		decoder    = decoder.New()
		auth       = auth.NewApiKey(apiKey)
	)
	return client.New(httpClient, decoder).
		WithAuth(auth)
}
