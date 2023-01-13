package lol

type SpectatorMatch struct {
	GameID            int64                        `json:"gameId"`
	MapID             int                          `json:"mapId"`
	GameMode          string                       `json:"gameMode"`
	GameType          string                       `json:"gameType"`
	GameQueueConfigID int                          `json:"gameQueueConfigId"`
	Participants      []SpectatorMatchParticipants `json:"participants"`
	Observers         SpectatorMatchObservers      `json:"observers"`
	PlatformID        string                       `json:"platformId"`
	BannedChampions   []interface{}                `json:"bannedChampions"`
	GameStartTime     int                          `json:"gameStartTime"`
	GameLength        int                          `json:"gameLength"`
}

type SpectatorMatchPerks struct {
	PerkIds      []int `json:"perkIds"`
	PerkStyle    int   `json:"perkStyle"`
	PerkSubStyle int   `json:"perkSubStyle"`
}

type SpectatorMatchParticipants struct {
	TeamID                   int                 `json:"teamId"`
	Spell1ID                 int                 `json:"spell1Id"`
	Spell2ID                 int                 `json:"spell2Id"`
	ChampionID               int                 `json:"championId"`
	ProfileIconID            int                 `json:"profileIconId"`
	SummonerName             string              `json:"summonerName"`
	Bot                      bool                `json:"bot"`
	SummonerID               string              `json:"summonerId"`
	GameCustomizationObjects []interface{}       `json:"gameCustomizationObjects"`
	Perks                    SpectatorMatchPerks `json:"perks"`
}

type SpectatorMatchObservers struct {
	EncryptionKey string `json:"encryptionKey"`
}
