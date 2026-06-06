package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"taskflow/app/facades"
)

type M20260604135339CreateTasksTable struct{}

// Signature The unique signature for the migration.
func (r *M20260604135339CreateTasksTable) Signature() string {
	return "20260604135339_create_tasks_table"
}

func (r *M20260604135339CreateTasksTable) Up() error {
	if !facades.Schema().HasTable("tasks") {
		return facades.Schema().Create("tasks", func(table schema.Blueprint) {
			table.ID()
			table.BigInteger("project_id")
			table.String("title")
			table.Text("description").Nullable()
			table.TimestampTz("start_date").Nullable()
			table.TimestampTz("end_date").Nullable()
			table.SmallInteger("status")
			table.SmallInteger("urgency")
			table.BigInteger("parent_task_id").Nullable()
			table.TimestampsTz()

			table.Foreign("project_id").References("id").On("projects").CascadeOnDelete()
			table.Foreign("parent_task_id").References("id").On("tasks").NullOnDelete()

			table.Index("project_id")
			table.Index("status")
			table.Index("parent_task_id")
		})
	}
	return nil
}

func (r *M20260604135339CreateTasksTable) Down() error {
	return facades.Schema().DropIfExists("tasks")
}
