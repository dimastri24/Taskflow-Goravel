package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"taskflow/app/facades"
)

type M20260604131310CreateUsersTable struct{}

// Signature The unique signature for the migration.
func (r *M20260604131310CreateUsersTable) Signature() string {
	return "20260604131310_create_users_table"
}

func (r *M20260604131310CreateUsersTable) Up() error {
	if !facades.Schema().HasTable("users") {
		return facades.Schema().Create("users", func(table schema.Blueprint) {
			table.ID()
			table.String("fullname")
			table.String("username")
			table.String("email")
			table.String("password")
			table.TimestampsTz()

			table.Unique("username")
			table.Unique("email")
		})
	}
	return nil
}

func (r *M20260604131310CreateUsersTable) Down() error {
	return facades.Schema().DropIfExists("users")
}
