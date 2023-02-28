package httpclient

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Response struct {
	message string
}

// To post message to all the channels - or boardcast
func BoardCastToAllChannelsAndServers(s *discordgo.Session, data []byte) {
	for _, guild := range s.State.Guilds {
		channels, _ := s.GuildChannels(guild.ID)
		for _, c := range channels {
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}
			s.ChannelMessageSend(
				c.ID,
				fmt.Sprintf("From DiscordConnector --  Message is %q", string(data)),
			)
		}
	}
}

// To post the message only to general channel
func SendMessageToGeneralChannel(s *discordgo.Session, data []byte) {
	for _, guild := range s.State.Guilds {
		channels, _ := s.GuildChannels(guild.ID)
		for _, c := range channels {
			// check if it's a type text or audio
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}
			if c.Name != "general" {
				continue
			}
			s.ChannelMessageSend(
				c.ID,
				fmt.Sprintf("From DiscordConnector --  Message is  %q", string(data)),
			)
		}
	}
}

// To post the message with guildid and channelId
func SendMessageToSpecificGuildAndChannel(s *discordgo.Session, guildId string, channelId string, data []byte) {
	channels, _ := s.GuildChannels(guildId)
	for _, c := range channels {
		if c.ID != channelId {
			continue
		}
		// check if it's a type text or audio
		if c.Type != discordgo.ChannelTypeGuildText {
			continue
		}
		s.ChannelMessageSend(
			c.ID,
			fmt.Sprintf("From DiscordConnector --  Message is %q", string(data)),
		)
	}
}
