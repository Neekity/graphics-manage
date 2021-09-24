package model

import (
	"go-project/graphics-manage/backend/common/helper"
	"gorm.io/gorm"
)

type (
	GraphicsMessageModel interface {
	}

	defaultGraphicsMessageModel struct {
		GormDB *gorm.DB
		table  string
	}

	GraphicsMessage struct {
		ID              uint          `json:"id" gorm:"primaryKey;column:id;type:int(10) unsigned auto_increment"`
		SenderUserID    int           `json:"sender_user_id" gorm:"column:sender_user_id;type:int(10) unsigned;not null;comment:消息发布者Id"`               // 消息发布者的id
		Type            string        `json:"type" gorm:"column:type;type:varchar(16);not null;comment:消息类型"`                                           // 消息类型
		Title           string        `json:"title" gorm:"column:title;type:varchar(256);not null;index:graphics_message_title_index;comment:标题"`       // 标题
		Abstract        string        `json:"abstract" gorm:"column:abstract;type:varchar(256);not null;comment:摘要"`                                    // 摘要
		CoverPictureURL string        `json:"cover_picture_url" gorm:"column:cover_picture_url;type:varchar(256);not null;comment:封面图"`                 // 封面图
		ChannelID       uint          `json:"channel_id" gorm:"column:channel_id;type:int(10) unsigned;not null;comment:频道id"`                          // 频道id
		Receivers       string        `json:"receivers" gorm:"column:receivers;type:text;not null;comment:接收者"`                                         // 接收者
		Content         string        `json:"content" gorm:"column:content;type:longtext;not null;comment:内容"`                                          // 内容
		Status          uint32        `json:"status" gorm:"column:status;type:tinyint(3) unsigned;not null;default:0;comment:消息状态：0-待发布中，1-发布中，2-发布完成"` // 消息状态：0-待发布中，1-发布中，2-发布完成
		SendTime        helper.MyTime `json:"send_time" gorm:"column:send_time;type:datetime;comment:预期发送时间"`                                           // 预期发送时间
		SendAt          helper.MyTime `json:"send_at" gorm:"column:send_at;type:datetime;comment:发送时间"`                                                 // 发送时间
		CreatedAt       helper.MyTime `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`       // 创建时间
		UpdatedAt       helper.MyTime `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`       // 更新时间
		Author          string        `json:"author" gorm:"column:author;type:varchar(16);comment:作者"`                                                  // 作者
	}
)

func NewGraphicsMessageModel(gdb *gorm.DB) GraphicsMessageModel {
	gdb.Set("gorm:table_options", "ENGINE=InnoDB COMMENT '图文消息表'  charset = utf8mb4;").
		AutoMigrate(&GraphicsMessage{})
	return &defaultGraphicsMessageModel{
		GormDB: gdb,
		table:  "`graphics_message`",
	}
}
