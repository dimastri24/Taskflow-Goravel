package models

import (
	"encoding/json"
	"time"

	"github.com/goravel/framework/database/orm"
)

type Task struct {
	orm.Model
	ProjectID    uint       `gorm:"column:project_id" json:"project_id"`
	Title        string     `gorm:"column:title" json:"title"`
	Description  *string    `gorm:"column:description" json:"description"` // Pointer for nullable text
	StartDate    *time.Time `gorm:"column:start_date" json:"start_date"`
	EndDate      *time.Time `gorm:"column:end_date" json:"end_date"`
	Status       int16      `gorm:"column:status" json:"status"`
	Urgency      int16      `gorm:"column:urgency" json:"urgency"`
	ParentTaskID *uint      `gorm:"column:parent_task_id" json:"parent_task_id"`
	Project      *Project   `gorm:"foreignKey:ProjectID" json:"project,omitempty"`
	ParentTask   *Task      `gorm:"foreignKey:ParentTaskID" json:"parent_task,omitempty"`
	SubTasks     []Task     `gorm:"foreignKey:ParentTaskID" json:"sub_tasks,omitempty"`
	// Assignees    []User        `gorm:"many2many:task_assigns;joinForeignKey:TaskID;joinReferences:UserID" json:"assignees,omitempty"`
	Comments    []TaskComment `gorm:"foreignKey:TaskID" json:"comments,omitempty"`
	Outputs     []TaskOutput  `gorm:"foreignKey:TaskID" json:"outputs,omitempty"`
	TaskAssigns []TaskAssign  `gorm:"foreignKey:TaskID" json:"-"`
}

// MarshalJSON otomatis mengubah struktur Task saat di-convert menjadi JSON API
func (t Task) MarshalJSON() ([]byte, error) {
	// Definisikan struktur penampung untuk "from" dan "to"
	type AssigneesGroup struct {
		From []User `json:"from"`
		To   []User `json:"to"`
	}

	// Alias digunakan agar tidak terjadi infinite loop saat marshalling
	type Alias Task

	// Inisialisasi slice kosong agar jika datanya tidak ada, di JSON tampil [] bukan null
	group := AssigneesGroup{
		From: make([]User, 0),
		To:   make([]User, 0),
	}

	// Lakukan perulangan untuk memisahkan User berdasarkan tipe-nya
	for _, assign := range t.TaskAssigns {
		switch assign.Type {
		case 1:
			group.To = append(group.To, assign.User)
		case 2:
			group.From = append(group.From, assign.User)
		}
	}

	// Satukan kembali ke dalam satu object JSON khusus
	return json.Marshal(&struct {
		Alias
		Assignees AssigneesGroup `json:"assignees"`
	}{
		Alias:     (Alias)(t),
		Assignees: group,
	})
}
