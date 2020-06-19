package app

import (
	"github.com/YukiUmetsu/go-bookstore-users-api/controllers/ping"
	"github.com/YukiUmetsu/go-bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
