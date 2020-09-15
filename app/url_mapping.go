package app

import (
	"github.com/carloscubas/bookstore_user-api/controllers/ping"
	"github.com/carloscubas/bookstore_user-api/controllers/user"
)

func mapUrls() {
	route.GET("/ping", ping.Ping)

	route.GET("/users/:user_id", user.GetUser)
	route.POST("/users", user.CreateUsers)
	route.PUT("/users/:user_id", user.UpdateUser)
	route.PATCH("/users/:user_id", user.UpdateUser)
}
