package entity

type Profile struct {
	ID             int64  `gorm:"primaryKey:auto_increment" json:"-"`
	Company        string `gorm:"type:varchar(100)" json:"-"`
	Website        string `gorm:"type:varchar(100)" json:"-"`
	Location       string `gorm:"type:varchar(100)" json:"-"`
	Status         string `gorm:"type:varchar(100)" json:"-"`
	Skills         string `gorm:"type:varchar(100)" json:"-"`
	GithubUsername string `gorm:"type:varchar(100)" json:"-"`
	Bio            string `gorm:"type:varchar(500)" json:"-"`
	Twitter        string `gorm:"type:varchar(100)" json:"-"`
	Facebook       string `gorm:"type:varchar(100)" json:"-"`
	Linkedin       string `gorm:"type:varchar(100)" json:"-"`
	Youtube        string `gorm:"type:varchar(100)" json:"-"`
	Instagram      string `gorm:"type:varchar(100)" json:"-"`
	UserID         string `gorm:"not null" json:"-"`
}
