package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"github.com/kosha/discord-connector/pkg/config"
	"github.com/kosha/discord-connector/pkg/logger"
)

type App struct {
	Router *mux.Router
	Log    logger.Logger
	Cfg    *config.Config
}

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}

func startBot() {
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session.
	dg.Close()
}

// Initialize creates the necessary scaffolding of the app
func (a *App) Initialize(log logger.Logger) {

	// Create a new Discord session using the provided bot token.
	go startBot()
	cfg := config.Get()

	a.Cfg = cfg
	a.Log = log

	a.Router = router()

	a.initializeRoutes()
}

// Run starts the app and serves on the specified addr
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!discord" {
		response, err := http.Get("https://github.com")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "./images/discord.png", response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get data from https://github.com")
		}
	}
}
