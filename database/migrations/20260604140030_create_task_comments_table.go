package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"taskflow/app/facades"
)

type M20260604140030CreateTaskCommentsTable struct{}

// Signature The unique signature for the migration.
func (r *M20260604140030CreateTaskCommentsTable) Signature() string {
	return "20260604140030_create_task_comments_table"
}

func (r *M20260604140030CreateTaskCommentsTable) Up() error {
	if !facades.Schema().HasTable("task_comments") {
		return facades.Schema().Create("task_comments", func(table schema.Blueprint) {
			table.ID()
			table.BigInteger("user_id")
			table.BigInteger("task_id")
			table.Text("message")
			table.BigInteger("parent_comment_id").Nullable()
			table.TimestampsTz()

			table.Foreign("user_id").References("id").On("users").CascadeOnDelete()
			table.Foreign("task_id").References("id").On("tasks").CascadeOnDelete()
			table.Foreign("parent_comment_id").References("id").On("task_comments").CascadeOnDelete()

			table.Index("task_id")
			table.Index("parent_comment_id")
		})
	}
	return nil
}

func (r *M20260604140030CreateTaskCommentsTable) Down() error {
	return facades.Schema().DropIfExists("task_comments")
}
