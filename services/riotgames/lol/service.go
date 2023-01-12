package lol

const (
	BaseUrl     = "na1.api.riotgames.com"
	SpectatorV4 = "/lol/spectator/v4/active-games/by-summoner/{encryptedSummonerId}"
)

type Client interface {
}

type CustomClient struct {
}

func NewCustomClient() *CustomClient {
	return &CustomClient{}
}
