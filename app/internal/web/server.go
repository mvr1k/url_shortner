package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"url_shortner/app/config"
)

type Server struct {
	server *gin.Engine
	apis   map[string]API //module name to API mapping
	logger *zap.SugaredLogger
}

func NewServer(logger *zap.SugaredLogger) *Server {
	ser := &Server{
		gin.Default(),
		make(map[string]API),
		logger,
	}
	ser.initialize()
	return ser
}

/*
initialize will do all initialization  related stuff
that is creating db instances and
creating module instances and
*/
func (server *Server) initialize() {
	//create cache instances

	//create db instance

	//create the modules here

	//assign modules to the server

	//create all routes here
}

/*addNewAPI map the module name to its manager*/
func (server *Server) addNewAPI(moduleName string, controller API) {
	server.apis[moduleName] = controller
}

/*ShutDownWithGrace in case of exit this will handle shutdown event with GRACE :)
 */
func (server *Server) ShutDownWithGrace() {
	server.logger.Infof("Good Bye Cruel World :( ")
}

func (server *Server) Start() error {
	address := fmt.Sprintf("%v:%v", config.ServerHost, config.ServerPort)

	return server.server.Run(address)
}
