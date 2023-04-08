# Kosha Discord Connector

Discord is a VoIP and instant messaging social platform. Users have the ability to communicate with voice calls, video calls, text messaging, media and files in private chats or as part of communities called "servers".

The connector APIs allow you to perform 'RESTful' operations such as reading, modifying, adding or deleting data from your helpdesk. The APIs also support Cross-Origin Resource Sharing (CORS).

In order to use Discord connector, you need to provide the bot access to a server for which you have admin access. Once the bot is added it can perform various operations as specified by the API specification.


## Bot vs User Accounts

Discord's API provides a separate type of user account dedicated to automation, called a bot account. Bot accounts can be created through the applications page, and are authenticated using a token (rather than a username and password). Unlike the normal OAuth2 flow, bot accounts have full access to most API routes without using bearer tokens, and can connect to the Real Time Gateway. Automating normal user accounts (generally called "self-bots") outside of the OAuth2/bot API is forbidden, and can result in account termination if found.

Bot accounts have a few differences in comparison to normal user accounts, namely:

* Bots are added to guilds through the OAuth2 API, and cannot accept normal invites.
* Bots cannot have friends, nor be added to or join Group DMs.
* Verified bots do not have a maximum number of Guilds.
* Bots have an entirely separate set of Rate Limits.

Discord connector uses a bot behind the scenes.

## What is possible with Discord APIs?

The Discord APIs provide your applications with direct access to the Discord Platform, giving you the ability to:

* Interact with channels
* Manage channel messages
* Guilds

and much more!

Discord's API is based around two core layers, a HTTPS/REST API for general operations, and persistent secure WebSocket based connection for sending and subscribing to real-time events. The most common use case of the Discord API will be providing a service, or access to a platform through the OAuth2 API.


## Useful Actions 

### Channels

Represents a guild or DM channel within Discord.

- Get channel information
- Get channel messages
- Create messages
- Delete message

### Guilds

Guilds in Discord represent an isolated collection of users and channels, and are often referred to as "servers"

- Get Guilds
- Modify Guilds
- Delete Guilds
- Create Guilds

Refer to the Discord connector [API specification](openapi.json) for details.

## Example Usage

< sdk example? >

## Authentication

Authenticating with the Discord API can be done in one of two ways:

- Using a bot token gained by registering a bot, for more information on bots see bots vs user accounts.
- Using an OAuth2 bearer token gained through the OAuth2 API.

For all authentication types, authentication is performed with the Authorization HTTP header in the format Authorization: TOKEN_TYPE TOKEN.