package controllers

import (
	"fmt"

	config "github.com/ankur12345678/shout/Config"
	"github.com/go-redis/redis/v8"
	gonanoid "github.com/matoous/go-nanoid/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BaseController struct {
	DB     *gorm.DB
	Config *config.Creds
	RedisClient *redis.Client
}

var Ctrl BaseController  //GLOBAL DB instance for controllers only

func UUIDGen(category string) string {
	id, err := gonanoid.New(10)
	if err != nil {
		log.Error("Error in generating uuid ", err)
	}
	if category == "COMMENT" {
		return fmt.Sprintf("c_%s", id)
	} else if category == "POST" {
		return fmt.Sprintf("p_%s", id)
	} else if category == "USER" {
		return fmt.Sprintf("u_%s", id)
	}
	return ""
}
