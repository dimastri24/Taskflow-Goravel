package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"taskflow/app/facades"
)

type M20260604140109CreateTaskOutputsTable struct{}

// Signature The unique signature for the migration.
func (r *M20260604140109CreateTaskOutputsTable) Signature() string {
	return "20260604140109_create_task_outputs_table"
}

func (r *M20260604140109CreateTaskOutputsTable) Up() error {
	if !facades.Schema().HasTable("task_outputs") {
		return facades.Schema().Create("task_outputs", func(table schema.Blueprint) {
			table.ID()
			table.BigInteger("task_id")
			table.SmallInteger("type")
			table.String("file_name").Nullable()
			table.String("file_ext").Nullable()
			table.Integer("file_size").Nullable()
			table.Text("link").Nullable()
			table.TimestampsTz()

			table.Foreign("task_id").References("id").On("tasks").CascadeOnDelete()

			table.Index("task_id")
		})
	}
	return nil
}

func (r *M20260604140109CreateTaskOutputsTable) Down() error {
	return facades.Schema().DropIfExists("task_outputs")
}
