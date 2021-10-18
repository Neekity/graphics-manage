package model

import (
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/tal-tech/go-zero/core/fx"
	"go-project/graphics-manage/backend/common/constant"
	"go-project/graphics-manage/backend/common/helper"
	"gorm.io/gorm"
	"strconv"
)

type (
	GraphicsMessageModel interface {
		Create(message GraphicsMessage) (*GraphicsMessage, error)
		List(title string, channelId int) ([]GraphicsMessage, error)
		FilterMessage(enforcer *casbin.Enforcer, messages []GraphicsMessage, userId string, dataType string) ([]GraphicsMessage, error)
		OwnerDetail(id int) (*MessageOwnerDetail, error)
		UserDetail(id int) (*MessageUserDetail, error)
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
		ChannelID       int           `json:"channel_id" gorm:"column:channel_id;type:int(10) unsigned;not null;comment:频道id"`                          // 频道id
		Receivers       string        `json:"receivers" gorm:"column:receivers;type:text;not null;comment:接收者"`                                         // 接收者
		Content         string        `json:"content" gorm:"column:content;type:longtext;not null;comment:内容"`                                          // 内容
		Status          uint32        `json:"status" gorm:"column:status;type:tinyint(3) unsigned;not null;default:0;comment:消息状态：0-待发布中，1-发布中，2-发布完成"` // 消息状态：0-待发布中，1-发布中，2-发布完成
		SendTime        helper.MyTime `json:"send_time" gorm:"column:send_time;type:datetime;comment:预期发送时间"`                                           // 预期发送时间
		SendAt          helper.MyTime `json:"send_at" gorm:"column:send_at;type:datetime;comment:发送时间"`                                                 // 发送时间
		CreatedAt       helper.MyTime `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`       // 创建时间
		UpdatedAt       helper.MyTime `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`       // 更新时间
		Author          string        `json:"author" gorm:"column:author;type:varchar(16);comment:作者"`                                                  // 作者
	}

	MessageOwnerDetail struct {
		CoverPictureURL string     `json:"cover_picture_url"`
		Title           string     `json:"title"`
		Abstract        string     `json:"abstract"`
		Content         string     `json:"content"`
		SubTitle        string     `json:"sub_title"`
		SenderUserName  string     `json:"sender_user_name"`
		ChannelName     string     `json:"channel_name"`
		SendAt          string     `json:"send_at"`
		Status          string     `json:"status"`
		Receivers       []Receiver `json:"receivers"`
	}

	MessageUserDetail struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		SubTitle string `json:"sub_title"`
	}

	Receiver struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
)

func (m *defaultGraphicsMessageModel) OwnerDetail(id int) (*MessageOwnerDetail, error) {
	message, err := m.findOne(id)
	if err != nil {
		return nil, err
	}
	var receivers []Receiver
	if err := json.Unmarshal([]byte(message.Receivers), &receivers); err != nil {
		return nil, err
	}
	var channel GraphicsChannel
	if err := m.GormDB.Model(&GraphicsChannel{}).First(&channel, message.ChannelID).Error; err != nil {
		return nil, err
	}
	var user User
	if err := m.GormDB.Model(&User{}).First(&user, message.SenderUserID).Error; err != nil {
		return nil, err
	}
	sendAt := message.SendAt.Format("2006-01-02")
	subTitle := sendAt
	if len(message.Author) != 0 {
		subTitle = message.Author + " " + sendAt
	}
	messageOwnerDetail := MessageOwnerDetail{
		CoverPictureURL: message.CoverPictureURL,
		Title:           message.Title,
		Abstract:        message.Abstract,
		Content:         message.Content,
		SubTitle:        subTitle,
		SenderUserName:  user.Name,
		ChannelName:     channel.Name,
		SendAt:          sendAt,
		Status:          constant.MessageStatus[int(message.Status)],
		Receivers:       receivers,
	}
	return &messageOwnerDetail, nil
}

func (m *defaultGraphicsMessageModel) UserDetail(id int) (*MessageUserDetail, error) {
	panic("implement me")
}

func (m *defaultGraphicsMessageModel) findOne(id int) (*GraphicsMessage, error) {
	var message GraphicsMessage
	if err := m.GormDB.First(&message, id).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (m *defaultGraphicsMessageModel) List(title string, channelId int) ([]GraphicsMessage, error) {
	var messages []GraphicsMessage
	query := m.GormDB.Table(m.table)
	if title != "" {
		query.Scopes(helper.QueryKey("%"+title+"%", "title", "like"))
	}
	if channelId != 0 {
		query.Scopes(helper.QueryKey(channelId, "channel_id", "="))
	}
	if err := query.Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (m *defaultGraphicsMessageModel) FilterMessage(enforcer *casbin.Enforcer, messages []GraphicsMessage, userId string, dataType string) ([]GraphicsMessage, error) {
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	var results []GraphicsMessage
	fx.From(func(source chan<- interface{}) {
		for _, message := range messages {
			source <- message
		}
	}).ForEach(func(material interface{}) {
		temp := material.(GraphicsMessage)
		if flag, _ := enforcer.Enforce(userId, fmt.Sprintf(constant.CasbinMessagePolicy, temp.ID), dataType); flag {
			results = append(results, temp)
		}
	})

	return results, nil
}

func (m *defaultGraphicsMessageModel) Create(message GraphicsMessage) (*GraphicsMessage, error) {
	err := m.GormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&message).Error; err != nil {
			return err
		}
		var receivers []Receiver
		if err := json.Unmarshal([]byte(message.Receivers), &receivers); err != nil {
			return err
		}
		var casbinModels []GraphicsCasbinRule
		for _, receiver := range receivers {
			casbinModels = append(casbinModels, GraphicsCasbinRule{
				Ptype: constant.CasbinPolicy,
				V0:    strconv.Itoa(receiver.Id),
				V1:    fmt.Sprintf(constant.CasbinMessagePolicy, message.ID),
				V2:    constant.CasbinPermissionRead,
			})
		}
		casbinModels = append(casbinModels, GraphicsCasbinRule{
			Ptype: constant.CasbinResourceRole,
			V0:    fmt.Sprintf(constant.CasbinMessagePolicy, message.ID),
			V1:    fmt.Sprintf(constant.CasbinChannelMessageResourceRole, message.ChannelID),
		})
		if err := tx.Create(&casbinModels).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func NewGraphicsMessageModel(gdb *gorm.DB) GraphicsMessageModel {
	gdb.Set("gorm:table_options", "ENGINE=InnoDB COMMENT '图文消息表'  charset = utf8mb4;").
		AutoMigrate(&GraphicsMessage{})
	return &defaultGraphicsMessageModel{
		GormDB: gdb,
		table:  "`graphics_message`",
	}
}
