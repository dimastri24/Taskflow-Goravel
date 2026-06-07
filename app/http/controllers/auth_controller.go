package controllers

import (
	"taskflow/app/facades"
	"taskflow/app/models"

	"github.com/goravel/framework/contracts/http"
)

type AuthController struct {
	// Dependent services can be injected here
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (r *AuthController) Register(ctx http.Context) http.Response {
	// 1. Define and Run Validation
	validator, err := ctx.Request().Validate(map[string]string{
		"fullname": "required",
		"username": "required",
		"email":    "required|email",
		"password": "required|min:6", // Adding min:6 for security best practices
	})

	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	if validator.Fails() {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"message": "Validation errors",
			"errors":  validator.Errors().All(),
		})
	}

	// 2. Extract Data from Form Request
	fullname := ctx.Request().Input("fullname")
	username := ctx.Request().Input("username")
	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")

	// 3. Check for unique constraints manually (or you could use 'unique:users,email' rule if configured)
	var existingUser models.User
	facades.Orm().Query().Where("username = ? OR email = ?", username, email).First(&existingUser)
	if existingUser.ID > 0 {
		return ctx.Response().Json(http.StatusConflict, http.Json{
			"message": "Username or Email already exists",
		})
	}

	// 4. Hash the password using Goravel's Hash facade
	hashedPassword, err := facades.Hash().Make(password)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "Failed to process password",
		})
	}

	// 5. Create and save the new User
	newUser := models.User{
		Fullname: fullname,
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	if err := facades.Orm().Query().Create(&newUser); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "Could not create user",
			"error":   err.Error(),
		})
	}

	// 6. Return successful response (Password is automatically omitted due to json:"-" tag)
	return ctx.Response().Json(http.StatusCreated, http.Json{
		"message": "User registered successfully",
		"data":    newUser,
	})
}
