// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameArchivedViewHistory = "archived_view_history"

// ArchivedViewHistory 浏览历史记录
type ArchivedViewHistory struct {
	Bvid       string         `gorm:"column:bvid;type:varchar(64);primaryKey" json:"bvid"`
	CreateTime *time.Time     `gorm:"column:create_time;type:datetime(3);autoCreateTime" json:"create_time"`
	UpdateTime *time.Time     `gorm:"column:update_time;type:datetime(3);autoUpdateTime" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:datetime(3)" json:"delete_time"`
	Resp       *string        `gorm:"column:resp;type:json" json:"resp"`
	VideoInfo  ArchivedVideo  `gorm:"foreignKey:bvid" json:"video_info"`
}

// TableName ArchivedViewHistory's table name
func (*ArchivedViewHistory) TableName() string {
	return TableNameArchivedViewHistory
}
