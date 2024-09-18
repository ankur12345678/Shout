package routes

import (
	"fmt"

	config "github.com/ankur12345678/shout/Config"
	controllers "github.com/ankur12345678/shout/Controllers"
	"github.com/ankur12345678/shout/Controllers/middlewares"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitRoutes(ctrl *controllers.BaseController) {
	log.Info("Initializing Routes .....")
	config := config.LoadConfig()
	r := gin.Default()
	r.GET("/", ctrl.RootHandler)
	r.POST("/signup", ctrl.SignUpHandler)
	r.POST("/login", ctrl.LoginHandler)
	r.POST("/blog", middlewares.HandleAuth, ctrl.InsertBlogHandler)
	r.PUT("/blog", middlewares.HandleAuth, ctrl.UpdateBlogHandler)
	r.GET("/blog/:id", middlewares.HandleAuth, ctrl.ShowBlogById)
	r.GET("/blog", middlewares.HandleAuth, ctrl.ShowAllBlogs)
	r.POST("/refresh", middlewares.HandleAuth, ctrl.HandleRefresh)

	r.POST("/logout", middlewares.HandleAuth, ctrl.HandleLogOut)
	log.Info("Initializing Routes : Success.....")
	fmt.Printf(`%s
     _______. __    __    ______    __    __  .___________.
    /       ||  |  |  |  /  __  \  |  |  |  | |           |
   |   (---- |  |__|  | |  |  |  | |  |  |  |  ---|  |----
    \   \    |   __   | |  |  |  | |  |  |  |     |  |     
.----)   |   |  |  |  | |   --   | |   --   |     |  |     
|_______/    |__|  |__|  \______/   \______/      |__|     
                                                           
`, "")
	r.Run(fmt.Sprintf(":%v", config.SERVER_PORT))
}
