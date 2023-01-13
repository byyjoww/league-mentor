package config

type Config struct {
	Environment string    `mapstructure:"environment"`
	Logging     Logging   `mapstructure:"logging"`
	Http        HTTP      `mapstructure:"http"`
	ChatGPT     ChatGPT   `mapstructure:"chatGpt"`
	RiotGames   RiotGames `mapstructure:"riotGames"`
}

type Logging struct {
	Level string `mapstructure:"level"`
}

type HTTP struct {
	Api       HTTPServer `mapstructure:"api"`
	Telemetry HTTPServer `mapstructure:"telemetry"`
}

type HTTPServer struct {
	Address string `mapstructure:"address"`
	Auth    Auth   `mapstructure:"auth"`
}

type Auth struct {
	Enabled bool   `mapstructure:"enabled"`
	User    string `mapstructure:"user"`
	Pass    string `mapstructure:"pass"`
}

type ChatGPT struct {
	ApiKey string `mapstructure:"apiKey"`
}

type RiotGames struct {
	ApiKey       string `mapstructure:"apiKey"`
	SummonerName string `mapstructure:"summonerName"`
}
