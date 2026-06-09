package controllers

import (
	"taskflow/app/facades"
	"taskflow/app/models"
	"taskflow/app/utils"

	"github.com/goravel/framework/contracts/http"
)

type WorkspaceController struct{}

func NewWorkspaceController() *WorkspaceController {
	return &WorkspaceController{}
}

func (r *WorkspaceController) Index(ctx http.Context) http.Response {
	// 1. Fetch authenticated user
	userRef, ok := ctx.Value("auth_user").(*models.User)
	if !ok || userRef == nil {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"message": "Unauthorized",
		})
	}

	var results []models.WorkspaceWithCounts

	// 2. Query workspaces where the owner matches currently logged-in user ID
	err := facades.Orm().Query().
		Model(&models.Workspace{}).
		Select(`
			workspaces.*, 
			(SELECT COUNT(*) FROM projects WHERE projects.workspace_id = workspaces.id) as projects_count,
			(SELECT COUNT(*) FROM workspace_members WHERE workspace_members.workspace_id = workspaces.id) as members_count
		`).
		Where("workspaces.owner = ?", userRef.ID).
		With("UserOwner").
		Find(&results)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "Failed to retrieve workspaces",
			"error":   err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Workspaces retrieved successfully",
		"data":    results,
	})
}

func (r *WorkspaceController) Store(ctx http.Context) http.Response {
	// 1. Define and validate the incoming request payload
	var req struct {
		Name string `form:"name" json:"name"`
	}

	rules := map[string]string{
		"name": "required|min_len:3|max_len:255",
	}

	if res := utils.BindAndValidate(ctx, &req, rules); res != nil {
		return res // Jika gagal, otomatis mengembalikan response error JSON
	}

	// 2. Retrieve the authenticated user from the context
	// Goravel's http.Context has a Value() method just like standard context
	userRef, ok := ctx.Value("auth_user").(*models.User)
	if !ok || userRef == nil {
		return ctx.Response().Json(http.StatusUnauthorized, http.Json{
			"message": "Unauthorized: User session not found",
		})
	}

	// 3. Prepare the model data
	workspace := models.Workspace{
		Name:  req.Name,
		Owner: userRef.ID, // Assuming your User model uses standard uint ID via orm.Model
	}

	// 4. Persist to database using Goravel's ORM Facade
	if err := facades.Orm().Query().Create(&workspace); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "Failed to create workspace",
			"error":   err.Error(),
		})
	}

	// 5. Return success response
	return ctx.Response().Json(http.StatusCreated, http.Json{
		"message": "Workspace created successfully",
		"data":    workspace,
	})
}
