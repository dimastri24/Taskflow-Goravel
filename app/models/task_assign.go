package models

type TaskAssign struct {
	TaskID uint  `gorm:"column:task_id;primaryKey"`
	UserID uint  `gorm:"column:user_id;primaryKey"`
	Type   int16 `gorm:"column:type;primaryKey"` // 1 = Assign To, 2 = Assign From
	User   User  `gorm:"foreignKey:UserID"`      // Mengambil profil data user
}
