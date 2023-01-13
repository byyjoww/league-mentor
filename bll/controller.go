package bll

import (
	"context"
	"fmt"

	"github.com/byyjoww/league-mentor/config"
	"github.com/byyjoww/league-mentor/services/chatgpt"
	"github.com/byyjoww/league-mentor/services/riotgames"
	"github.com/byyjoww/league-mentor/services/riotgames/lol"
)

type Controller interface {
	GetIdentity(ctx context.Context, summonerName string) (lol.Summoner, error)
	GetCurrentMatchOverview(ctx context.Context, puuid string) (string, error)
}

type GenericController struct {
	gptClient    *chatgpt.GoGptClient
	leagueClient *riotgames.Client
}

func NewGenericController(cfg config.Config) *GenericController {
	return &GenericController{
		gptClient:    chatgpt.NewGoGptClient(cfg.ChatGPT),
		leagueClient: riotgames.NewClient(cfg.RiotGames),
	}
}

func (c *GenericController) GetIdentity(ctx context.Context, summonerName string) (lol.Summoner, error) {
	return c.leagueClient.LeagueOfLegends.GetSummonerByName(ctx, summonerName)
}

func (c *GenericController) GetCurrentMatchOverview(ctx context.Context, puuid string) (string, error) {
	matches, err := c.leagueClient.LeagueOfLegends.GetMatchesByPuuid(ctx, puuid)
	if err != nil {
		return "", err
	}

	if len(matches) < 1 {
		return "", err
	}

	match, err := c.leagueClient.LeagueOfLegends.GetMatchById(ctx, matches[0])
	if err != nil {
		return "", err
	}

	self, err := match.GetSelf(puuid)
	if err != nil {
		return "", err
	}

	opponent, err := match.GetOpponent(self)
	if err != nil {
		return "", err
	}

	prompt := createPrompt(self.Lane, self.ChampionName, opponent.ChampionName)

	completion, err := c.gptClient.CreateCompletion(ctx, prompt)
	if err != nil {
		return "", err
	}

	return completion.Value, nil
}

func createPrompt(lane string, championSelf string, championOpponent string) string {
	prompt := fmt.Sprintf("How do I win against %s as %s in %s?", championOpponent, championSelf, lane)
	// fmt.Println("Prompt:", prompt)
	return prompt
}
