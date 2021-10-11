// Code generated by goctl. DO NOT EDIT.
package types

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ApiListData struct {
	CurPage    int         `json:"curPage"`
	LastPage   int         `json:"lastPage"`
	List       interface{} `json:"list"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
}

type StoreChannelRequest struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Audiences []Audience `json:"audiences"`
}

type ChanenlDetailRequest struct {
	Id int `json:"id"`
}

type Audience struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SearchChannelRequest struct {
	Name string `json:"name"`
}

type StoreMaterialRequest struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Url       string `json:"url"`
	Content   string `json:"content"`
	Abstract  string `json:"abstract"`
	ChannelId int    `json:"channel_id"`
}

type MaterialImageUploadRequest struct {
	ChannelId int `json:"channel_id"`
}

type DeleteMaterialRequest struct {
	Id uint `json:"id"`
}

type MaterialDetailRequest struct {
	Id int `json:"id"`
}

type SearchMaterialRequest struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	ChannelId int    `json:"channel_id"`
}

type PublishMaterialRequest struct {
	Id        int      `json:"id"`
	Receivers []string `json:"receivers"`
	SendTime  string   `json:"send_time"`
	Author    string   `json:"author"`
}

type DeleteMenuRequest struct {
	Id uint `json:"id"`
}

type StoreMenuRequest struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parent_id"`
	Path     string `json:"path"`
	Icon     string `json:"icon"`
}

type MessageChangeStatusRequest struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}

type SearchMessageRequest struct {
	Name      string `json:"name"`
	Page      int    `json:"page"`
	Type      string `json:"type"`
	ChannelId int    `json:"channel_id"`
}

type OaUserInfo struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type JwtTokenResponse struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"` // 建议客户端刷新token的绝对时间
}

type BasicUserInfo struct {
	Id       int    `json:"id"`
	Userid   string `json:"userid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Position string `json:"position"`
}

type UserInfo struct {
	Code    int           `json:"code;"`
	Data    BasicUserInfo `json:"data"`
	Message string        `json:"message"`
}

type SearchUserRequest struct {
	Name string `json:"name"`
}

type AssignRolerRequest struct {
	Id          uint     `json:"id"`
	CasbinRoles []string `json:"casbin_roles"`
}

type UserRolesRequest struct {
	UserId uint `json:"user_id"`
}

type SearchRoleRequest struct {
	Name string `json:"name"`
}

type StoreAndAssignRolerRequest struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	CasbinRole string `json:"casbin_role"`
	UserIds    []uint `json:"user_ids"`
}

type RoleDetailRequest struct {
	Id uint `json:"id"`
}
