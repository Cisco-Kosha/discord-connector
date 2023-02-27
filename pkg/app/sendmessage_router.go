package app

import (
	"net/http"

	"github.com/kosha/freshservice-connector/pkg/httpclient"
)

// boardCastToAllChannelsAndServers godoc
// @Summary Get details of the connector specification
// @Description Send messages to all channels
// @Tags Messages
// @Accept  json
// @Produce  json
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

// sendMessageToAllGeneralChannel godoc
// @Summary Send message to general Channel
// @Description Send messages to General Channel
// @Tags Messages
// @Accept  json
// @Produce  json
// @Success 200 {object} object
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v1/boardcast/general [post]
func (a *App) sendMessageToAllGeneralChannel(w http.ResponseWriter, r *http.Request) {
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
	httpclient.SendMessageToGeneralChannel(dg)
	respondWithJSON(w, http.StatusOK, map[string]string{
		"Data Sent": "Data Sent to General Channels in a Server",
	})
}
