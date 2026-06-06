package bootstrap

import (
	"github.com/goravel/framework/contracts/database/schema"

	"taskflow/database/migrations"
)

func Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20210101000001CreateJobsTable{},
		&migrations.M20260604131310CreateUsersTable{},
		&migrations.M20260604134927CreateWorkspacesTable{},
		&migrations.M20260604135112CreateWorkspaceMembersTable{},
		&migrations.M20260604135306CreateProjectsTable{},
		&migrations.M20260604135339CreateTasksTable{},
		&migrations.M20260604135915CreateTaskAssignsTable{},
		&migrations.M20260604140030CreateTaskCommentsTable{},
		&migrations.M20260604140109CreateTaskOutputsTable{},
	}
}
