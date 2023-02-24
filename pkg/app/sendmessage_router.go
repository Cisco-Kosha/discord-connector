package app

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/kosha/freshservice-connector/pkg/httpclient"
)

func (a *App) boardCastToAllChannelsAndServers(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	dg, err := a.Cfg.GetDiscordSession()
	// Open a websocket connection to Discord and begin listening.
	// Discord Bot is online here
	err = dg.Open()
	// Close a websocket connection as soon as the api request is complete
	// Discord Bot is offline here
	defer dg.Close()
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
	}
	// make a httpclient api call
	httpclient.BoardCastToAllChannelsAndServers(dg)
	respondWithJSON(w, http.StatusOK, map[string]string{
		"Data Sent": "Data Sent to all channels",
	})
}

func (a *App) sendMessageToAllGeneralChannel(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	dg, err := discordgo.New(a.Cfg.GetBotToken())
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	// Open a websocket connection to Discord and begin listening.
	// Discord Bot is online here
	err = dg.Open()
	// Close a websocket connection as soon as the api request is complete
	// Discord Bot is offline here
	defer dg.Close()
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
	}
	// make a httpclient api call
	httpclient.SendMessageToGeneralChannel(dg)
	respondWithJSON(w, http.StatusOK, map[string]string{
		"Data Sent": "Data Sent to General Channels in a Server",
	})
}
