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
	workspaceController := controllers.NewWorkspaceController()

	// Route group for auth
	facades.Route().Prefix("auth").Group(func(router route.Router) {
		router.Post("/register", authController.Register)
		router.Post("/login", authController.Login)

		router.Middleware(middleware.JwtAuth()).Group(func(authRouter route.Router) {
			authRouter.Post("/logout", authController.Logout)
			// POST	/refresh
		})
	})

	facades.Route().Middleware(middleware.JwtAuth()).Group(func(router route.Router) {
		router.Get("/profile", userController.Profile)
		// PUT	  /profile

		router.Get("/workspaces", workspaceController.Index)
		router.Post("/workspaces", workspaceController.Store)
		// GET    /workspaces/{id}
		// PUT    /workspaces/{id}
		// DELETE /workspaces/{id}

		// GET	  /workspaces/{id}/members
		// POST	  /workspaces/{id}/members

		// GET    /workspaces/{id}/projects
		// POST   /workspaces/{id}/projects

		// GET    /projects
		// GET    /projects/{id}

		// GET    /projects/{id}/task
		// POST   /projects/{id}/task

		// GET    /tasks
		// GET    /tasks/{id}
		// PUT    /tasks/{id}
		// DELETE /tasks/{id}

		// GET    /tasks/{id}/comments
		// POST   /tasks/{id}/comments
	})
}
