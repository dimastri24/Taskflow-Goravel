package routes

import (
	"taskflow/app/facades"
	"taskflow/app/http/controllers"

	"github.com/goravel/framework/contracts/route"
)

func Api() {
	authController := controllers.NewAuthController()

	// Route group for auth
	facades.Route().Prefix("auth").Group(func(router route.Router) {
		router.Post("/register", authController.Register)
	})
}
