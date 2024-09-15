package controllers

import (
	config "github.com/ankur12345678/shout/Config"
	"gorm.io/gorm"
)

type BaseController struct {
	DB     *gorm.DB
	Config *config.Creds
}
