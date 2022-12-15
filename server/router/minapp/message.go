package minapp

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MessageRouter struct{}

func (r *MessageRouter) InitMessageRouter(Router *gin.RouterGroup) {
	messagePublicRouter := Router.Group("message")
	messagePrivateRouter := Router.Group("message").Use(middleware.JWTAuthMiddleware())
	messageApi := v1.ApiGroupApp.MinAppApiGroup.MessageApi
	{
		messagePublicRouter.GET("second-online", messageApi.SecondOnlineMessage)
		messagePublicRouter.GET("count", messageApi.MessageNotify)

	}

	{
		//创建二手闲置私信房间
		messagePrivateRouter.POST("second-room", messageApi.CreateSecondRoom)

		//获取二手闲置私信房间信息及消息内容
		messagePrivateRouter.GET("second-details", messageApi.GetSecondRoomMessage)

		//获取私信列表
		messagePrivateRouter.GET("online", messageApi.GetOnlineMessage)

		//获取系统消息列表
		messagePrivateRouter.GET("system", messageApi.GetSystemMessage)

		//获取留言列表
		messagePrivateRouter.GET("leave", messageApi.GetLeaveMessage)

		//获取未读消息数量
		messagePrivateRouter.GET("unread", messageApi.GetUnreadMessage)

		//删除留言信息
		messagePrivateRouter.DELETE(":id", messageApi.DeleteMessage)
	}
}
