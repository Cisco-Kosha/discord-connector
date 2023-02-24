package httpclient

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// To post message to all the channels - or boardcast
func BoardCastToAllChannelsAndServers(s *discordgo.Session) {
	for _, guild := range s.State.Guilds {
		channels, _ := s.GuildChannels(guild.ID)
		for _, c := range channels {
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}
			s.ChannelMessageSend(
				c.ID,
				fmt.Sprintf("From DiscordConnector --  Channel name is %q", c.Name),
			)
		}
	}
}

// To post the message only to general channel
func SendMessageToGeneralChannel(s *discordgo.Session) {
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
				fmt.Sprintf("From DiscordConnector --  Channel name is %q", c.Name),
			)
		}
	}
}
