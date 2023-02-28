package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	var apiV1 = "/api/v1"

	// specification routes
	a.Router.HandleFunc(apiV1+"/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	//  routes
	a.Router.HandleFunc(apiV1+"/broadcast", a.boardCastToAllChannelsAndServers).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/broadcast/general", a.sendMessageToAllGeneralChannel).Methods("POST", "OPTIONS")

	// swagger-docs
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)

}
