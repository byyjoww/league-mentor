# League Mentor

## Overview
League Mentor will provide you advice on how to play your League of Legends matches. 

## Running League Mentor

#### Running a Server
- Add your ChatGTP api key to config.yaml
- Add your Riot Games api key to config.yaml
- Run the terminal command `make run-api`

#### Running a Client
- Add your summoner name to config.yaml
- Run the terminal command `make run-client`

## Testing
If you want to build a server or client locally and not be worried about committing your api keys, you can create a file called `dev-env.json` in the `config` directory with the following json object:
```go
{
    "chatGptApiKey": "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",
    "riotGamesApiKey": "RGAPI-aa9ae8f8-8cbd-42ac-ba50-ad0368eb3a97"
}
```
You can then use the `make dev-run-api` or `make dev-run-client` commands to start a server/client while overriding the api key variables in `config.yaml` with the variables set in `dev-env.json`.


## Roadmap
1. Update client to connect to server via tcp
2. Add general advice on how to behave in lane once the game starts
3. Add dynamic advice throughout the match on how to behave based on the game state (ahead or behind your opponent)
4. Add macro advice on how to close out the game based on the match progress