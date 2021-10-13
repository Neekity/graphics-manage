package constant

/**
 * 消息状态：发布中
 */
const MessageStatusPublishing = 0

/**
 * 消息状态：发布完成
 */
const MessageStatusFinished = 1

/**
 * 消息状态：已删除
 */
const MessageStatusDeleted = 2

/**
 * 消息状态：取消发布
 */
const MessageStatusCANCELED = 3

/**
 * 消息状态
 */
var MessageStatus = map[int]string{
	MessageStatusPublishing: "发布中",
	MessageStatusFinished:   "发布完成",
	MessageStatusDeleted:    "已删除",
	MessageStatusCANCELED:   "取消发布",
}

/**
 * 素材类型
 */
const MaterialImage = "image"

const MaterialGraphics = "graphics"

const MaterialTemplate = "template"

var MaterialMaps = map[string]string{
	MaterialImage:    "图片",
	MaterialGraphics: "图文",
	MaterialTemplate: "模版",
}

/**
 * casbin数据操作权限
 */
const CasbinPermissionWrite = "write"

const CasbinPermissionRead = "read"

/**
 * casbin资源角色
 */
const CasbinResourceRole = "g2"

const CasbinDefaultRole = "g"

const CasbinPolicy = "p"

/**
 * casbin权限连接符号
 */
const CasbinPermissionSymbol = "/"

const CasbinChannelWriteRole = "graphics/channel/%d/write"
const ChannelWriteRoleName = "%s频道写角色"
const CasbinChannelReadRole = "graphics/channel/%d/read"
const ChannelReadRoleName = "%s频道读角色"

const ChannelRole = "graphics/channel/%d"

const CasbinChannelMaterialResourceRole = "graphics/channel/%d/material"
const CasbinChannelMessageResourceRole = "graphics/channel/%d/message"

const CasbinMessagePolicy = "graphics/message/%d"
const CasbinMaterialPolicy = "graphics/material/%d"
