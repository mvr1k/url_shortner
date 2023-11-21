package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"url_shortner/app/config"
	"url_shortner/app/internal/shortner"
	"url_shortner/app/internal/web/api"
)

type Server struct {
	server *gin.Engine
	apis   map[string]api.API //module name to API mapping
	logger *zap.SugaredLogger
}

const version = "/api/v1"

func NewServer(logger *zap.SugaredLogger) *Server {
	ser := &Server{
		gin.Default(),
		make(map[string]api.API),
		logger,
	}
	ser.initialize()
	return ser
}

func (server *Server) Start() error {
	address := fmt.Sprintf("%v:%v", config.ServerHost, config.ServerPort)

	return server.server.Run(address)
}

/*ShutDownWithGrace in case of exit this will handle shutdown event with GRACE :)*/
func (server *Server) ShutDownWithGrace() {
	server.logger.Infof("Good Bye Cruel World :( ")
}

/*
initialize will do all initialization  related stuff
that is creating db instances and
creating module instances and
*/
func (server *Server) initialize() {
	//create cache instances if required

	//create db instance if requried

	//create the modules here
	shortnerModule := shortner.NewWebApi(server.logger)
	server.register(shortnerModule)
	//assign modules to the server

	//create all routes here
}

/*addNewAPI map the module name to its manager*/
func (server *Server) register(apiController api.API) {
	//storing api controller address
	server.apis[apiController.ModuleName()] = apiController

	//adding all routes
	apiMappings := apiController.GetRouteMapping()
	for route, method := range apiMappings {
		server.addRoute(method, route, apiController)
	}
	server.logger.Infof("added %v module to server", apiController.ModuleName())

}

// addRoute will add a route in gin.Engine
func (server *Server) addRoute(route string, method string, module api.API) {
	server.server.Handle(method, version+route, api.GetMappedMethod(method, module))
}
