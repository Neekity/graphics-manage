# graphics-manage

类似企业微信图文消息和微信公众号图文消息的管理平台

# 后端

1. go-zero框架
2. casbin权限控制
3. gorm关系模型

# 前端

1. vue2
2. vuetify
3. tinymce

# dev环境一键安装

`./init.sh dev && docker-compose up -d`
tips：如果是amd64架构需要将go.Dockerfile中的GOARCH设置为amd64

#OAuth登陆说明
1. 可以接入已有的SSO系统
2. 可以使用https://www.authing.cn/网站

# 与手机APP交互说明
1. 可对接微信
2. 可对接telegram，由于telegram不支持图文卡片，发送带标题的url即可
3. 如有自己开发的app，可使用极光推送将消息推送至app

# 效果图如下
[点击进入demo👈](https://graphics.huangtaohome.top/)

<img src="https://i.ibb.co/jw2CvxJ/271639627860-pic-hd.png" alt="271639627860-pic-hd" border="0">
<img src="https://i.ibb.co/Gcq3JmG/281639627889-pic-hd.png" alt="281639627889-pic-hd" border="0">

# QA
Q：为啥把前端的dist文件夹也给加进来了
A：因为服务器太垃圾无法build，只能在本地build
