# Kosha Discord Connector

Discord is a social platform that enables you to communicate via voice, video, or chat—and to share files and media—in public or private communities called "guilds" (or "servers").

The Kosha Discord connector enables you to perform REST API operations from the Discord API in your Kosha workflow or custom application. Using the Kosha Discord connector, you can directly access the Discord platform to:

* Interact with channels
* Manage channel messages
* Mange guilds

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

## Authentication

 When provisioning the Kosha Discord connector, authentication to the Discord API can be done by bot or OAuth2.

### Bot Authentication

In Discord, bots are added to all applications and are authenticated using a token, as opposed to a username or password. To use a bot for authentication with the Kosha Discord connector, you must add the bot to a server for which you have admin access. After you add the bot and grant it access, it can perform any operations provided in the API specification. For details about bots, refer to the [Discord Bot documentation](https://discord.com/developers/docs/topics/oauth2#bot-users).

### OAuth2 Authentication

Use an OAuth2 bearer token gained through the [Discord OAuth2 API](https://discord.com/developers/docs/topics/oauth2#oauth2).

## Kosha Connector Open Source Development

All connectors Kosha shares on the marketplace are open source. We believe in fostering collaboration and open development. Everyone is welcome to contribute their ideas, improvements, and feedback for any Kosha connector. We encourage community engagement and appreciate any contributions that align with our goals of an open and collaborative API management platform.

Refer to the contribution guidelines for details.

## Contributing

Pull requests and bug reports are welcome.

For larger changes, please create an issue in GitHub first to discuss your proposed changes and their possible implications.