package controllers

import (
	"fmt"
	"taskflow/app/facades"
	"taskflow/app/models"
	"taskflow/app/utils"

	"github.com/goravel/framework/contracts/http"
)

type WorkspaceController struct{}

func NewWorkspaceController() *WorkspaceController {
	return &WorkspaceController{}
}

// Index handles: GET /workspaces
func (r *WorkspaceController) Index(ctx http.Context) http.Response {
	// 1. Fetch authenticated user
	userRef, errUser := utils.GetUser(ctx)
	if errUser != nil {
		return errUser
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

// Store handles: POST /workspaces
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
	userRef, errUser := utils.GetUser(ctx)
	if errUser != nil {
		return errUser
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

// Show handles: GET /workspaces/{id}
func (r *WorkspaceController) Show(ctx http.Context) http.Response {
	// Fetch authenticated user
	userRef, errUser := utils.GetUser(ctx)
	if errUser != nil {
		return errUser
	}

	// Get ID from the route parameter
	id := ctx.Request().Route("id")

	var result models.WorkspaceWithCounts

	// Fetch workspace only if it belongs to the logged-in user
	err := facades.Orm().Query().
		Model(&models.Workspace{}).
		Select(`
			workspaces.*, 
			(SELECT COUNT(*) FROM projects WHERE projects.workspace_id = workspaces.id) as projects_count,
			(SELECT COUNT(*) FROM workspace_members WHERE workspace_members.workspace_id = workspaces.id) as members_count
		`).
		Where("id = ? AND owner = ?", id, userRef.ID).
		With("UserOwner").
		FirstOrFail(&result) // Use First to get a single record

	if err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "Workspace not found or access denied",
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Workspace retrieved successfully",
		"data":    result,
	})
}

// Update handles: PUT /workspaces/{id}
func (r *WorkspaceController) Update(ctx http.Context) http.Response {
	userRef, errUser := utils.GetUser(ctx)
	if errUser != nil {
		return errUser
	}

	id := ctx.Request().Route("id")

	fmt.Println("route id:", id)
	fmt.Println("user:", userRef.ID)

	// 1. Verify workspace existence and ownership first
	var workspace models.Workspace
	err := facades.Orm().Query().Where("id = ? AND owner = ?", id, userRef.ID).FirstOrFail(&workspace)
	if err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "Workspace not found or access denied",
		})
	}

	fmt.Printf("%+v\n", workspace)
	fmt.Println("ID:", workspace.ID)
	fmt.Printf("err: %#v\n", err)

	// 2. Bind and Validate incoming data
	var req struct {
		Name string `form:"name" json:"name"`
	}
	rules := map[string]string{
		"name": "required|min_len:3|max_len:255",
	}
	if res := utils.BindAndValidate(ctx, &req, rules); res != nil {
		return res
	}

	// 3. Update data
	workspace.Name = req.Name
	if err := facades.Orm().Query().Save(&workspace); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "Failed to update workspace",
			"error":   err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Workspace updated successfully",
		"data":    workspace,
	})
}

// Destroy handles: DELETE /workspaces/{id}
func (r *WorkspaceController) Destroy(ctx http.Context) http.Response {
	userRef, errUser := utils.GetUser(ctx)
	if errUser != nil {
		return errUser
	}

	id := ctx.Request().Route("id")

	// Find the workspace ensuring scope strictly checks ownership
	var workspace models.Workspace
	err := facades.Orm().Query().Where("id = ? AND owner = ?", id, userRef.ID).FirstOrFail(&workspace)
	if err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "Workspace not found or access denied",
		})
	}

	// Delete record from database
	if _, err := facades.Orm().Query().Delete(&workspace); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "Failed to delete workspace",
			"error":   err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Workspace deleted successfully",
	})
}
