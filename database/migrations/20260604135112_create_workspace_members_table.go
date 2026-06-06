package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"taskflow/app/facades"
)

type M20260604135112CreateWorkspaceMembersTable struct{}

// Signature The unique signature for the migration.
func (r *M20260604135112CreateWorkspaceMembersTable) Signature() string {
	return "20260604135112_create_workspace_members_table"
}

func (r *M20260604135112CreateWorkspaceMembersTable) Up() error {
	if !facades.Schema().HasTable("workspace_members") {
		return facades.Schema().Create("workspace_members", func(table schema.Blueprint) {
			table.ID()
			table.BigInteger("user_id")
			table.BigInteger("workspace_id")
			table.TimestampTz("joined_at") // Menggunakan format dengan Timezone

			table.Foreign("user_id").References("id").On("users").CascadeOnDelete()
			table.Foreign("workspace_id").References("id").On("workspaces").CascadeOnDelete()

			table.Unique("user_id", "workspace_id")
		})
	}
	return nil
}

func (r *M20260604135112CreateWorkspaceMembersTable) Down() error {
	return facades.Schema().DropIfExists("workspace_members")
}
