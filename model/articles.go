package model

type ArticlesModel struct {
	BaseModel
	Title    string `gorm:"column:title; not null" json:"title"`
	Content  string `gorm:"column:content" json:"content"`
	Category string `gorm:"column:category" json:"category"`
	Status   int8   `gorm:"column:status; not null" json:"status"`
}
