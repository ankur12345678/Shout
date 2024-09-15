package main

import (
	"fmt"

	config "github.com/ankur12345678/shout/Config"
	controllers "github.com/ankur12345678/shout/Controllers"
	migration "github.com/ankur12345678/shout/Migration"
	routes "github.com/ankur12345678/shout/Routes"
	log "github.com/sirupsen/logrus"
)

func main() {
	//defining routes
	log.Info("Starting Server...")
	fmt.Printf(`%s
     _______. __    __    ______    __    __  .___________.
    /       ||  |  |  |  /  __  \  |  |  |  | |           |
   |   (---- |  |__|  | |  |  |  | |  |  |  |  ---|  |----
    \   \    |   __   | |  |  |  | |  |  |  |     |  |     
.----)   |   |  |  |  | |   --   | |   --   |     |  |     
|_______/    |__|  |__|  \______/   \______/      |__|     
                                                           
`, "")


	//loading config
	config := config.LoadConfig()
	db := migration.InitDB()
	ctrl := controllers.BaseController{
		DB:     db,
		Config: config,
	}
	routes.InitRoutes(&ctrl)
}
