package app

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kosha/discord-connector/pkg/httpclient"
)

// boardCastToAllChannelsAndServers godoc
// @Summary Get details of the connector specification
// @Description Send messages to all channels
// @Tags Messages
// @Accept  json
// @Produce  json
// @Param Message body object false "Enter Message to be Sent to Discord"
// @Success 200 {object} object
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v1/broadcast [post]
func (a *App) boardCastToAllChannelsAndServers(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var data []byte
	if r.Body != nil {
		data, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}
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
	httpclient.BoardCastToAllChannelsAndServers(dg, data)
	respondWithJSON(w, http.StatusOK, map[string]string{
		"Data Sent": "Data Sent to all channels",
	})
}

// sendMessageToAllGeneralChannel godoc
// @Summary Send message to general Channel
// @Description Send messages to General Channel
// @Tags Messages
// @Accept  json
// @Produce  json
// @Success 200 {object} object
// @Param Message body object false "Enter Message to be Sent to Discord"
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v1/broadcast/general [post]
func (a *App) sendMessageToAllGeneralChannel(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var data []byte
	if r.Body != nil {
		data, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}
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
	httpclient.SendMessageToGeneralChannel(dg, data)
	respondWithJSON(w, http.StatusOK, map[string]string{
		"Data Sent": "Data Sent to General Channels in a Server",
	})
}

// sendMessageToSpecificGuildAndChannel godoc
// @Summary Send message to general Channel
// @Description Send messages to General Channel
// @Tags Messages
// @Accept  json
// @Produce  json
// @Param channelId path string true "Channel Id"
// @Param guildId path string true "Guild Id"
// @Param Message body object false "Enter Message to be Sent to Discord"
// @Success 200 {object} object
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v1/guild/{guildId}/channel/{channelId} [post]
func (a *App) sendMessageToSpecificGuildAndChannel(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var data []byte
	if r.Body != nil {
		data, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}
	vars := mux.Vars(r)
	guildId := vars["guildId"]
	channelId := vars["channelId"]
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
	httpclient.SendMessageToSpecificGuildAndChannel(dg, guildId, channelId, data)
	respondWithJSON(w, http.StatusOK, map[string]string{
		"Data Sent": "Data Sent to General Channels in a Server",
	})
}
