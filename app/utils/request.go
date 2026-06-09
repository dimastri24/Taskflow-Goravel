package utils

import (
	"net/http"

	ghttp "github.com/goravel/framework/contracts/http"
)

// BindAndValidate handles binding, validation, and automatically returns the error response if it fails.
// If it returns nil, it means validation passed and you can safely use your struct data.
func BindAndValidate(ctx ghttp.Context, req interface{}, rules map[string]string) ghttp.Response {
	// 1. Bind the incoming data
	if err := ctx.Request().Bind(req); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, ghttp.Json{
			"message": "Invalid request format",
			"error":   err.Error(),
		})
	}

	// 2. Validate the request
	validator, err := ctx.Request().Validate(rules)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, ghttp.Json{
			"message": "Validation system error",
			"error":   err.Error(),
		})
	}

	// 3. Check if validation failed
	if validator.Fails() {
		return ctx.Response().Json(http.StatusUnprocessableEntity, ghttp.Json{
			"message": "Validation failed",
			"errors":  validator.Errors().All(),
		})
	}

	// Success! Return nil so the controller knows it can proceed.
	return nil
}
