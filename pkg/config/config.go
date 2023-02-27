package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	botToken   string
	domainName string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.botToken, "bottoken", os.Getenv("BOT_TOKEN"), "Discord BOT TOKEN")
	flag.Parse()

	return conf
}

func (c *Config) GetBotToken() string {
	token := "Bot " + c.botToken
	return token
}

func (c *Config) GetDomainName() string {
	return c.domainName
}

func (c *Config) GetDiscordSession() (s *discordgo.Session, err error) {
	dg, err := discordgo.New(c.GetBotToken())
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	return dg, err
}
