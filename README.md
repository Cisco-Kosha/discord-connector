# Kosha Discord Connector

![Discord](images/Discord.png)

Discord is a social platform that enables you to communicate via voice, video, or chat—and to share files and media—in public or private communities called "guilds" (or "servers").

Using the Kosha Discord connector, you can perform REST API operations such as reading, modifying, adding, or deleting data from your Discord servers. 

Using the Discord API, your Kosha workflow or application can directly access the Discord platform to:

* Interact with channels
* Manage channel messages
* Mange guilds

Discord's API is based on two core layers, a HTTPS/REST API for general operations and a persistent secure WebSocket connection for sending and subscribing to real-time events. The Discord API is most commonly used to provide services, or access to a platform, through the OAuth2 API. 

## Useful Actions 

You can use the Kosha Discord connector to perform several useful operations to manage your Discord channels and guilds. 

Refer to the Kosha Discord connector [API specification](openapi.json) for details.

### Managing Channels

- Get channel information
- Get channel messages
- Create messages
- Delete message

### Managing Guilds

- Get Guilds
- Modify Guilds
- Delete Guilds
- Create Guilds

## Example Usage

< sdk example? >

## Authentication

Authentication to the Discord API can be done by bot or OAuth 2:

### Bot Authentication

In Discord, bots are added to all applications and are authenticated using a token, as opposed to a username or password. To use a bot for authentication with the Kosha Discord connector, you must add the bot to a server for which you have admin access. After you add the bot and grant it access, it can perform any operations provided in the API specification. For details about bots, refer to the [Discord Bot documentation](https://discord.com/developers/docs/topics/oauth2#bot-users).

### OAuth 2 Authentication

Use an OAuth2 bearer token gained through the [Discord OAuth2 API](https://discord.com/developers/docs/topics/oauth2#oauth2).

For all authentication types, authentication is performed with the Authorization HTTP header in the following format:

`Authorization: TOKEN_TYPE TOKEN`