package main

import (
	config "github.com/ankur12345678/shout/Config"
	controllers "github.com/ankur12345678/shout/Controllers"
	migration "github.com/ankur12345678/shout/Migration"
	routes "github.com/ankur12345678/shout/Routes"
	log "github.com/sirupsen/logrus"
)

func main() {
	//defining routes
	log.Info("Starting Server...")

	//loading config
	config := config.LoadConfig()
	db := migration.InitDB()
	ctrl := controllers.BaseController{
		DB:     db,
		Config: config,
	}
	controllers.Ctrl = ctrl
	migration.SeedDB(db)
	routes.InitRoutes(&ctrl)

}
