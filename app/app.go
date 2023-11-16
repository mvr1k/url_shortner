package app

import (
	"github.com/common-nighthawk/go-figure"
	"url_shortner/app/common/logger"
	"url_shortner/app/config"
	"url_shortner/app/internal/web"
)

// on start print a ascii art
func init() {
	heading := figure.NewFigure("Shotgun.IO", "", true)
	heading.Print()
}

func Start() {
	//initialize config
	config.InitializeConfigs()

	//importing logger package initializes logger and will
	LOGGER := logger.Logger()
	LOGGER.Infof("Initializing Server.......")

	server := web.NewServer(LOGGER)
	defer server.ShutDownWithGrace()

	err := server.Start()
	if err != nil {
		LOGGER.Errorf("Error In Shutting Down Server : %v", err)
		panic("Server Shutting Down ... ")
	}

}
