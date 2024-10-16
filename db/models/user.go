package models

type User struct {
	BaseModel
	Username string `gorm:"type:string;size:20;not null;unique"`
	Password string `gorm:"type:string;size:100;not null"`
	UserRole *[]UserRole
}

type Role struct {
	BaseModel
	Name     string `gorm:"type:string;size:10;not null,unique"`
	UserRole *[]UserRole
}

type UserRole struct {
	BaseModel
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId int
	RoleId int
}
