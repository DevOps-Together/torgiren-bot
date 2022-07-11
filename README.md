# discord-bot

Discord bot - autorole bot for discord.

## Setup 
1. Create discord app and a bot. Provide generated bot token either as `-token` parameter or as TOKEN env variable. 
1. In "Oauth2 -> URL generator" tab choose `bot` scope, add required priviledges:
    1. Manage roles
    1. Manage channels
    1. Send Messages
    1. Read Messages/View Channels
    1. Manage Messages
1. Use generated URL to add bot to your server.
1. Prepare config file and pass it's path as parameter or use `example/config.yaml`
1. Run `bin/discord-bot`

## Build application
Run `make`
