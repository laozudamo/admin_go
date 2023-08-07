package models

import "gorm.io/gorm"

type SysPost struct {
	PostName   string `json:"postName" gorm:"size:128;comment:岗位名称"` // 岗位名称
	PostCode   string `json:"postCode" gorm:"size:128;comment:岗位代码"` // 岗位代码
	PostStatus string `json:"postStatus" gorm:"size:4;comment:岗位状态"` // 岗位状态（0正常 1停用）
	Sort       int    `json:"sort" gorm:"size:4;comment:岗位排序"`       // 岗位排序
	Remark     string `json:"remark" gorm:"size:255;comment:备注"`     // 备注
	gorm.Model
}
