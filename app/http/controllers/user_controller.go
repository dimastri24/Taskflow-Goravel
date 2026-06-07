package controllers

import (
	"taskflow/app/models"

	"github.com/goravel/framework/contracts/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// func (r *UserController) Index(ctx http.Context) http.Response {
// 	return ctx.Response().Success().Json(http.Json{
// 		"Hello": "Goravel",
// 	})
// }

func (r *UserController) Profile(ctx http.Context) http.Response {
	// Ambil pointer user dari middleware secara aman
	if userRef, ok := ctx.Value("auth_user").(*models.User); ok {
		return ctx.Response().Json(http.StatusOK, http.Json{
			"message": "Profile retrieved successfully",
			"data":    userRef, // Password otomatis di-omit oleh tag json:"-" di model
		})
	}

	// Fallback aman jika karena suatu hal data di context kosong (tidak akan panic)
	return ctx.Response().Json(http.StatusUnauthorized, http.Json{
		"message": "Unauthorized: User context not found",
	})
}
