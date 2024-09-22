package migration

import (
	models "github.com/ankur12345678/shout/Models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SeedDB(DB *gorm.DB) {
	userRepo := models.InitUserRepo(DB)
	postRepo := models.InitPostRepo(DB)
	err := userRepo.Create(&models.User{UserUUID: "user_1", Name: "Ankur", Email: "a@gmail.com", UserName: "ankur11", Password: "hehe", ProfilePicture: "yo.svg"})
	if err != nil {
		log.Error("err in creating user entry ", err)
	}

	err = postRepo.Create(&models.Post{PostUUID: "post_1", Title: "Welcome to GOLang", Content: "hi everyone!", UserID: 1, Likes: 78, Replies: 56})
	if err != nil {
		log.Error("err in creating post entry", err)
	}

	log.Info("-----SEEDING SUCCESS-----")
}
