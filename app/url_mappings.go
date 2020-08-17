package app

import (
	"github.com/sug5806/bookstore_users-api/controllers/ping"
	"github.com/sug5806/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.FindUser)
	router.POST("/users", users.CreateUser)

}
