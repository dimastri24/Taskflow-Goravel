package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"taskflow/app/facades"
)

type M20260604135915CreateTaskAssignsTable struct{}

// Signature The unique signature for the migration.
func (r *M20260604135915CreateTaskAssignsTable) Signature() string {
	return "20260604135915_create_task_assigns_table"
}

func (r *M20260604135915CreateTaskAssignsTable) Up() error {
	if !facades.Schema().HasTable("task_assigns") {
		return facades.Schema().Create("task_assigns", func(table schema.Blueprint) {
			table.BigInteger("task_id")
			table.BigInteger("user_id")
			table.SmallInteger("type")

			table.Foreign("task_id").References("id").On("tasks").CascadeOnDelete()
			table.Foreign("user_id").References("id").On("users").CascadeOnDelete()

			table.Primary("task_id", "user_id", "type")
		})
	}
	return nil
}

func (r *M20260604135915CreateTaskAssignsTable) Down() error {
	return facades.Schema().DropIfExists("task_assigns")
}
