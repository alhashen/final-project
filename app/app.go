package app

import (
	"finalproj/config"
	"finalproj/repository"
	"finalproj/route"
	"finalproj/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.GORM.DB)
	app := service.NewService(repo)
	route.RegisterAPI(router, app)

	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
