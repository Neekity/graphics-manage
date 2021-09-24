package model

import (
	"go-project/graphics-manage/backend/common/helper"
	"gorm.io/gorm"
)

type (
	GraphicsMessageReceiverModel interface {
	}

	defaultGraphicsMessageReceiverModel struct {
		GormDB *gorm.DB
		table  string
	}

	GraphicsMessageReceiver struct {
		ID                     uint          `json:"id" gorm:"primaryKey;column:id;type:int(10) unsigned auto_increment"`
		MessageID              uint          `json:"message_id" gorm:"column:message_id;type:int(10) unsigned;not null;index:graphics_message_receiver_message_id_index;comment:信息id"`                                            // 信息id
		ReceiverEmployeeNumber string        `json:"receiver_employee_number" gorm:"column:receiver_employee_number;type:varchar(16);not null;index:graphics_message_receiver_receiver_employee_number_index;comment:接收消息者的员工编号"` // 接收消息者的员工编号
		ReceiverEmployeeName   string        `json:"receiver_employee_name" gorm:"column:receiver_employee_name;type:varchar(16);not null;comment:接收消息者的姓名"`                                                                      // 接收消息者的姓名
		CreatedAt              helper.MyTime `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`                                                                          // 创建时间
		UpdatedAt              helper.MyTime `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`                                                                          // 更新时间
	}
)

func NewGraphicsMessageReceiverModel(gdb *gorm.DB) GraphicsMessageReceiverModel {
	gdb.Set("gorm:table_options", "ENGINE=InnoDB COMMENT '图文消息接收者表'  charset = utf8mb4;").
		AutoMigrate(&GraphicsMessageReceiver{})
	return &defaultGraphicsMessageReceiverModel{
		GormDB: gdb,
		table:  "`graphics_message_receiver`",
	}
}
