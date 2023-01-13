package riotgames

import (
	"github.com/byyjoww/league-mentor/config"
	"github.com/byyjoww/league-mentor/services/riotgames/lol"
)

type Client struct {
	LeagueOfLegends lol.Client
}

func NewClient(cfg config.RiotGames) *Client {
	return &Client{
		LeagueOfLegends: lol.NewCustomClient(cfg),
	}
}
