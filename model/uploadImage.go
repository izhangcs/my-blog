package model

type UploadImagesModel struct {
	BaseModel
	Path string `gorm:"column:path;not null" json:"path"`
	Type string `gorm:"column:type; not null" json:"type"`
	Size uint32 `gorm:"column:size; not null" json:"size"`
}
