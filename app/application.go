package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	if err := router.Run(":8888"); err != nil {
		panic(err)
	}
}
