package models

type Section struct {
	BaseModel
	Name string `gorm:"type:string;size:200;not null;unique"`
	Post []Post
}

type Post struct {
	BaseModel
	User        User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId      int
	Title       string `gorm:"type:string;size:200"`
	Description string
	Section     Section `gorm:"foreignKey:SectionId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	SectionId   int
}

type Image struct {
	BaseModel
	Path string `gorm:"type:string;size:200;not null;unique"`
}
