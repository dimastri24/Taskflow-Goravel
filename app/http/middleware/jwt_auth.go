package middleware

import (
	"taskflow/app/facades"
	"taskflow/app/models"

	"github.com/goravel/framework/contracts/http"
)

func JwtAuth() http.Middleware {
	return func(ctx http.Context) {
		// 1. Ambil token dari header
		token := ctx.Request().Header("Authorization")
		if token == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "Token is missing",
			})
			return
		}

		// 2. Parse token secara eksplisit (mengatasi isu otomatisasi caching Goravel)
		_, err := facades.Auth(ctx).Parse(token)
		if err != nil {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "Invalid or expired token",
				"error":   err.Error(),
			})
			return
		}

		// 3. Pastikan user-nya memang eksis di database kita
		var user models.User
		err = facades.Auth(ctx).User(&user)
		if err != nil || user.ID == 0 {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"message": "User not found or unauthorized",
			})
			return
		}

		// Simpan pointer user ke context
		ctx.WithValue("auth_user", &user)

		// Lolos semua pengecekan, lanjutkan ke Middleware berikutnya atau Controller
		ctx.Request().Next()
	}
}
