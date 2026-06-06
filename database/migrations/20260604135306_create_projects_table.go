package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"taskflow/app/facades"
)

type M20260604135306CreateProjectsTable struct{}

// Signature The unique signature for the migration.
func (r *M20260604135306CreateProjectsTable) Signature() string {
	return "20260604135306_create_projects_table"
}

func (r *M20260604135306CreateProjectsTable) Up() error {
	if !facades.Schema().HasTable("projects") {
		return facades.Schema().Create("projects", func(table schema.Blueprint) {
			table.ID()
			table.BigInteger("workspace_id")
			table.String("name")
			table.TimestampsTz()

			table.Foreign("workspace_id").References("id").On("workspaces").CascadeOnDelete()
		})
	}
	return nil
}

func (r *M20260604135306CreateProjectsTable) Down() error {
	return facades.Schema().DropIfExists("projects")
}
