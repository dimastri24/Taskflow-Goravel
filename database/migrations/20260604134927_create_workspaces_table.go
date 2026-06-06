package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"taskflow/app/facades"
)

type M20260604134927CreateWorkspacesTable struct{}

// Signature The unique signature for the migration.
func (r *M20260604134927CreateWorkspacesTable) Signature() string {
	return "20260604134927_create_workspaces_table"
}

func (r *M20260604134927CreateWorkspacesTable) Up() error {
	if !facades.Schema().HasTable("workspaces") {
		return facades.Schema().Create("workspaces", func(table schema.Blueprint) {
			table.ID()
			table.String("name")
			table.BigInteger("owner")
			table.TimestampsTz()

			table.Foreign("owner").References("id").On("users").CascadeOnDelete()
		})
	}
	return nil
}

func (r *M20260604134927CreateWorkspacesTable) Down() error {
	return facades.Schema().DropIfExists("workspaces")
}
