package routes

import (
	"taskflow/app/facades"
	"taskflow/app/http/controllers"
	"taskflow/app/http/middleware"

	"github.com/goravel/framework/contracts/route"
)

func Api() {
	authController := controllers.NewAuthController()
	userController := controllers.NewUserController()

	// Route group for auth
	facades.Route().Prefix("auth").Group(func(router route.Router) {
		router.Post("/register", authController.Register)
		router.Post("/login", authController.Login)

		router.Middleware(middleware.JwtAuth()).Group(func(authRouter route.Router) {
			authRouter.Post("/logout", authController.Logout) // <--- Tambahkan ini
		})
	})

	facades.Route().Middleware(middleware.JwtAuth()).Group(func(router route.Router) {
		// Endpoint Profile baru sesuai standar REST API
		router.Get("/profile", userController.Profile)

		// Kamu bisa tambah route terproteksi lainnya di dalam grup ini nanti, misal:
		// router.Put("/profile", userController.UpdateProfile)
	})
}
