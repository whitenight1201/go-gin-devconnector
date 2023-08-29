package entity

type User struct {
	ID       int64  `gorm:"primaryKey:auto_increment" json:"-"`
	Name     string `gorm:"type:varchar(100)" json:"-"`
	Email    string `gorm:"type:varchar(100)" json:"-"`
	Password string `gorm:"type:varchar(100)" json:"-"`
}
