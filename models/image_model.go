package models

import "fmt"

type ImageModel struct {
	Model
	UserID    uint      `gorm:"column:userID;comment:用户ID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	FileName  string    `gorm:"column:fileName;comment:文件名" json:"fileName"`
	Size      int64     `gorm:"column:size;comment:文件大小" json:"size"`
	// update/xx.png
	Path      string    `gorm:"column:path;comment:文件路径" json:"path"`
	Hash      string    `gorm:"column:hash;comment:文件的hash" json:"hash"`
}

// /update/xx.png
//相对于项目的 Web 根路径
func (image ImageModel) WebPath() string {
	return fmt.Sprintf("/%s", image.Path)
}
