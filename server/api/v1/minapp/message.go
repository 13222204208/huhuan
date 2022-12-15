package minapp

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/minapp"
	minAppRequest "github.com/flipped-aurora/gin-vue-admin/server/model/minapp/request"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type MessageApi struct{}

// ClientManager 客户端管理
type ClientManager struct {
	//客户端 map 储存并管理所有的长连接client，在线的为true，不在的为false
	clients map[*Client]bool
	//web端发送来的的message我们用broadcast来接收，并最后分发给所有的client
	broadcast chan []byte
	//谁发来的信息
	cid string
	//发来的信息人的房间号
	room string
	//新创建的长连接client
	register chan *Client
	//新注销的长连接client
	unregister chan *Client
}

// Client 客户端
type Client struct {
	//房间号
	room string
	//用户id
	id string
	//消息接收的id
	tid string
	//对应的订单号
	orderNum string
	//连接的socket
	socket *websocket.Conn
	//发送的消息
	send chan []byte
}

//会把Message格式化成json
type Message struct {
	//消息struct
	Sender    string `json:"sender,omitempty"`    //发送者
	Recipient string `json:"recipient,omitempty"` //接收者
	Content   string `json:"content,omitempty"`   //内容
}

//创建客户端管理者
var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
	cid:        "",
	room:       "",
}

//socket 设置
var (
	_ = websocket.Upgrader{
		//
		ReadBufferSize: 1023,
		//
		WriteBufferSize: 1023,
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (manager *ClientManager) start() {
	for {
		select {
		//如果有新的连接接入,就通过channel把连接传递给conn
		case conn := <-manager.register:
			//把客户端的连接设置为true
			manager.clients[conn] = true
			//把返回连接成功的消息json格式化
		//	jsonMessage, _ := json.Marshal(&Message{Content: "id:" + manager.cid + " 已连接"})
		//调用客户端的send方法，发送消息
		//	manager.send(jsonMessage, conn)
		//如果连接断开了
		case conn := <-manager.unregister:
			//判断连接的状态，如果是true,就关闭send，删除连接client的值
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				//	jsonMessage, _ := json.Marshal(&Message{Content: "id:" + manager.cid + " 已连接"})
				//	manager.send(jsonMessage, conn)
			}
			//广播
		case message := <-manager.broadcast:
			//遍历已经连接的客户端，把消息发送给他们
			for conn := range manager.clients {
				//只发送同一个房间号的
				//判断发送给谁（不发送自己）
				//接收者id： conn.id
				//发送者id： manager.cid
				if conn.room != manager.room || conn.id == manager.cid {
					continue
				}

				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

//定义客户端管理的send方法
func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		fmt.Println("send消息", string(message))
		//不给屏蔽的连接发送消息
		if conn != ignore {
			conn.send <- message
		}
	}
}

//定义客户端结构体的read方法
func (c *Client) read() {
	defer func() {
		//结构体cid赋值
		manager.cid = c.id
		manager.room = c.room
		//触发关闭
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		//读取消息
		_, message, err := c.socket.ReadMessage()
		//如果有错误信息，就注销这个连接然后关闭
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		//如果没有错误信息就把信息放入broadcast
		fmt.Println("broadcast消息", string(message))
		userId, _ := strconv.ParseInt(c.id, 10, 64)
		tid, _ := strconv.ParseInt(c.tid, 10, 64)
		if err := messageService.SaveSecondOnlineMessage(uint(userId), uint(tid), c.room, string(message), c.orderNum); err != nil {
			fmt.Println(err.Error())
		}

		//发送未读消息数量到首页
		_ = UnreadMessageCount([]string{c.tid})

		fmt.Println("发送的消息", string(message))
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		//结构体cid赋值
		manager.cid = c.id
		manager.room = c.room
		//触发消息发送
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		//从send里读消息
		case message, ok := <-c.send:
			//如果没有消息
			if !ok {
				_ = c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//有消息就写入，发送给web端
			_ = c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (a *MessageApi) SecondOnlineMessage(c *gin.Context) {
	//开一个goroutine执行开始程序
	go manager.start()

	wsHandler(c.Writer, c.Request)
}

func wsHandler(res http.ResponseWriter, req *http.Request) {
	//将http协议升级成websocket协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	//可以用传来的参数识别身份
	_ = req.ParseForm()
	room := req.Form["room"][0]
	uid := req.Form["uid"][0]
	tid := req.Form["tid"][0]
	orderNum := req.Form["orderNum"][0]

	//这里是随机生成id
	//每一次连接都会新开一个client，client.id通过uuid生成保证每次都是不同的
	//uuid.Must(uuid.NewV4()).String()
	client := &Client{id: uid, tid: tid, socket: conn, send: make(chan []byte), room: room, orderNum: orderNum}
	//注册一个新的链接
	manager.register <- client
	//启动协程收web端传过来的消息
	go client.read()
	//启动协程把消息返回给web端
	go client.write()
}

//创建 二手闲置的私信房间
func (a *MessageApi) CreateSecondRoom(c *gin.Context) {
	var room minapp.SecondRoom
	_ = c.ShouldBindJSON(&room)
	//判断是否传递当前私信的二手商品id
	if len(room.OrderNum) < 16 {
		response.FailWithMessage("请填写订单号", c)
		return
	}

	//判断是否传递卖家的用户ID
	if room.Sell < 1 {
		response.FailWithMessage("请填写卖家的用户ID", c)
		return
	}
	//当前用户id  有意思买家的ID
	uid := c.MustGet("id").(uint)
	room.Buy = uid

	_, roomId := messageService.SaveSecondRoom(room)
	response.OkWithDetailed(gin.H{"roomId": roomId}, "房间号返回成功", c)
}

//获取二手闲置房间聊天信息，及关联的商品信息
func (a *MessageApi) GetSecondRoomMessage(c *gin.Context) {
	var room minAppRequest.SecondRoomMessageRequest
	_ = c.ShouldBindQuery(&room)

	if room.RoomId < 1 {
		response.FailWithMessage("请填写正确的房间号", c)
		return
	}
	//获取用户的id为了使用读取的消息为已读
	uid := c.MustGet("id").(uint)
	if err, roomInfo, roomMessage := messageService.GetSecondRoomMessage(room.RoomId); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		if err := messageService.UpdateSecondRoomMessage(room.RoomId, uid); err != nil {
			response.FailWithMessage("修改消息为已读失败", c)
		} else {
			response.OkWithData(gin.H{"roomInfo": roomInfo, "roomMessage": roomMessage}, c)
		}

	}
}

//消息通知客户端
type SystemClient struct {
	ID            string          // 连接ID
	AccountId     string          // 账号id, 一个账号可能有多个连接
	Socket        *websocket.Conn // 连接
	HeartbeatTime int64           // 前一次心跳时间

}

//消息类型
const (
	MessageTypeHeartbeat = "heartbeat" // 心跳
	MessageTypeRegister  = "register"  // 注册

	HeartbeatCheckTime = 90  // 心跳检测几秒检测一次
	HeartbeatTime      = 200 // 心跳距离上一次的最大时间

	ChanBufferRegister   = 100 // 注册chan缓冲
	ChanBufferUnregister = 100 // 注销chan大小
)

// 客户端管理
type SystemClientManager struct {
	Clients  map[string]*SystemClient // 保存连接
	Accounts map[string][]string      // 账号和连接关系,map的key是账号id即：AccountId，这里主要考虑到一个账号多个连接
	mu       *sync.Mutex
}

// 定义一个管理Manager
var SystemManager = SystemClientManager{
	Clients:  make(map[string]*SystemClient), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Accounts: make(map[string][]string),      // 账号和连接关系
	mu:       new(sync.Mutex),
}

var (
	RegisterChan   = make(chan *SystemClient, ChanBufferRegister)   // 注册
	unregisterChan = make(chan *SystemClient, ChanBufferUnregister) // 注销
)

// 封装回复消息
type ServiceMessage struct {
	Type    string                `json:"type"` // 类型
	Content ServiceMessageContent `json:"content"`
}
type ServiceMessageContent struct {
	Body     string `json:"body"`      // 主要数据
	MetaData string `json:"meta_data"` // 扩展数据
}

func CreateReplyMsg(t string, content ServiceMessageContent) []byte {
	replyMsg := ServiceMessage{
		Type:    t,
		Content: content,
	}
	msg, _ := json.Marshal(replyMsg)
	return msg
}

// 注册注销
func register() {
	for {
		select {
		case conn := <-RegisterChan: // 新注册，新连接
			// 加入连接,进行管理
			accountBind(conn)
			fmt.Println("加入成功")
			// 回复消息
			content := CreateReplyMsg(MessageTypeRegister, ServiceMessageContent{})
			_ = conn.Socket.WriteMessage(websocket.TextMessage, content)

		case conn := <-unregisterChan: // 注销，或者没有心跳
			// 关闭连接
			_ = conn.Socket.Close()

			// 删除Client
			unAccountBind(conn)
		}
	}
}

// 绑定账号
func accountBind(c *SystemClient) {
	SystemManager.mu.Lock()
	defer SystemManager.mu.Unlock()

	// 加入到连接
	SystemManager.Clients[c.ID] = c

	// 加入到绑定
	if _, ok := SystemManager.Accounts[c.AccountId]; ok { // 该账号已经有绑定，就追加一个绑定
		SystemManager.Accounts[c.AccountId] = append(SystemManager.Accounts[c.AccountId], c.ID)
	} else { // 没有就新增一个账号的绑定切片
		SystemManager.Accounts[c.AccountId] = []string{c.ID}
	}
}

// 解绑账号
func unAccountBind(c *SystemClient) {
	SystemManager.mu.Lock()
	defer SystemManager.mu.Unlock()

	// 取消连接
	delete(SystemManager.Clients, c.ID)

	// 取消绑定
	if len(SystemManager.Accounts[c.AccountId]) > 0 {
		for k, clientId := range SystemManager.Accounts[c.AccountId] {
			if clientId == c.ID { // 找到绑定客户端Id
				SystemManager.Accounts[c.AccountId] = append(SystemManager.Accounts[c.AccountId][:k], SystemManager.Accounts[c.AccountId][k+1:]...)
			}
		}
	}
}

// 维持心跳
func heartbeat() {
	for {
		// 获取所有的Clients
		SystemManager.mu.Lock()
		clients := make([]*SystemClient, len(SystemManager.Clients))
		for _, c := range SystemManager.Clients {
			clients = append(clients, c)
		}
		SystemManager.mu.Unlock()

		for _, c := range clients {
			if time.Now().Unix()-c.HeartbeatTime > HeartbeatTime {
				unAccountBind(c)
			}
		}

		time.Sleep(time.Second * HeartbeatCheckTime)
	}
}

// 管理连接
func Start() {
	// 检查心跳
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		heartbeat()
	}()

	// 注册注销
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
			}
		}()
		register()
	}()
}

// 根据账号获取连接
func GetClient(accountId string) []*SystemClient {
	clients := make([]*SystemClient, 0)

	SystemManager.mu.Lock()
	defer SystemManager.mu.Unlock()
	fmt.Println("客户端长度", SystemManager.Accounts[accountId])
	if len(SystemManager.Accounts[accountId]) > 0 {
		for _, clientId := range SystemManager.Accounts[accountId] {
			if c, ok := SystemManager.Clients[clientId]; ok {
				clients = append(clients, c)
			}
		}
	}

	return clients
}

// 读取信息，即收到消息
func (c *SystemClient) SystemRead() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		// 读取消息
		_, body, err := c.Socket.ReadMessage()
		fmt.Println("读到的数据", body)
		if err != nil {
			break
		}

		var msg struct {
			Type string `json:"type"`
		}
		err = json.Unmarshal(body, &msg)

		if err != nil {
			log.Println(err)
			continue
		}

		if msg.Type == MessageTypeHeartbeat { // 维持心跳消息
			// 刷新连接时间
			c.HeartbeatTime = time.Now().Unix()

			// 回复心跳
			replyMsg := CreateReplyMsg(MessageTypeHeartbeat, ServiceMessageContent{})
			err = c.Socket.WriteMessage(websocket.TextMessage, replyMsg)
			if err != nil {
				log.Println(err)
			}
			continue
		}
	}
}

type Unread struct {
	Total int64 `json:"total"`
}

//首页消息数量统计
func UnreadMessageCount(s []string) error {
	//未读消息统计
	for _, uid := range s {
		r := messageService.GetUnreadMessageCount(uid)
		var read Unread
		read.Total = r
		msg, err := json.Marshal(read)
		if err != nil {
			return err
		}
		clients := GetClient(uid)
		for _, c := range clients {
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}

	return nil
}

// 发送消息
func SystemSend(accounts []string, message ServiceMessage) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}

	for _, accountId := range accounts {
		// 获取连接id
		clients := GetClient(accountId)
		//fmt.Println("获取客户端", clients)
		//fmt.Println("帐号", accountId, "消息内容", msg)
		// 发送消息
		for _, c := range clients {
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}

	return nil
}

type MessageNotifyRequest struct {
	UserId string `form:"user_id"`
}

func (a *MessageApi) MessageNotify(ctx *gin.Context) {
	Start()
	// 获取参数
	var params MessageNotifyRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		log.Println(err)
		return
	}
	// TODO: 鉴权

	// 将http升级为websocket
	conn, err := (&websocket.Upgrader{
		// 1. 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(ctx.Writer, ctx.Request, nil) // 升级
	if err != nil {
		log.Println(err)
		http.NotFound(ctx.Writer, ctx.Request)
		return
	}

	// 创建一个实例连接
	ConnId := xid.New().String()
	client := &SystemClient{
		ID:            ConnId, // 连接id
		AccountId:     fmt.Sprintf("%s", params.UserId),
		HeartbeatTime: time.Now().Unix(),
		Socket:        conn,
	}

	// 用户注册到用户连接管理
	RegisterChan <- client
	//clients := GetClient("2")
	//fmt.Println("获取的客户端", clients)
	// 读取信息
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("MessageNotify read panic: %+v", r)
			}
		}()

		client.SystemRead()
	}()
}

func SystemSendMessage(accountId []string, contents string) {

	var myMessage ServiceMessage
	var detailMessage ServiceMessageContent
	detailMessage.Body = contents
	detailMessage.MetaData = "消息meta"
	myMessage.Type = "系统消息"
	myMessage.Content = detailMessage
	//发送在线的客户端用户，用户消息内容
	err := SystemSend(accountId, myMessage)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("成功")
	}
}

func (a *MessageApi) GetSystemMessage(c *gin.Context) {
	uid := c.MustGet("id").(uint)

	var s autocode.SystemMessage
	_ = c.ShouldBindQuery(&s)
	if s.ID > 0 {
		if err, r := messageService.GetSystemMessageDetail(s.ID); err != nil {
			response.FailWithMessage("获取失败", c)
			return
		} else {
			response.OkWithData(gin.H{"data": r}, c)
			return
		}
	}

	if err, r := messageService.GetSystemMessage(uid); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": r}, c)
	}
}

//用户私信列表
func (a *MessageApi) GetOnlineMessage(c *gin.Context) {
	var room minapp.SecondRoom
	_ = c.ShouldBindQuery(&room)
	uid := c.MustGet("id").(uint)

	if room.ID > 0 {
		if err, r := messageService.GetRoomOnlineMessage(room.ID); err != nil {
			response.FailWithMessage("获取失败", c)
			return
		} else {

			if err := messageService.UpdateSecondRoomMessage(room.ID, uid); err != nil {
				response.FailWithMessage("修改消息为已读失败", c)
				return
			}
			response.OkWithData(gin.H{"list": r}, c)
			return
		}
	}

	if err, r := messageService.GetOnlineMessage(uid); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {

		response.OkWithData(gin.H{"list": r}, c)
	}
}

//用户留言列表
func (a *MessageApi) GetLeaveMessage(c *gin.Context) {
	var m minapp.SecondMessage
	_ = c.ShouldBindQuery(&m)
	uid := c.MustGet("id").(uint)
	if m.ID > 0 {
		if err, r := messageService.GetLeaveMessageDetail(m.ID); err != nil {
			response.FailWithMessage("获取失败", c)
			return
		} else {
			response.OkWithData(r, c)
			return
		}
	}

	if err, r := messageService.GetLeaveMessage(uid); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"list": r}, c)
	}
}

//获取未读消息的数量
func (a *MessageApi) GetUnreadMessage(c *gin.Context) {
	uid := c.MustGet("id").(uint)
	id := strconv.Itoa(int(uid))
	r := messageService.GetUnreadMessageCount(id)
	response.OkWithData(r, c)
}

//删除留言信息
func (a *MessageApi) DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	uid := c.MustGet("id").(uint)

	if err := messageService.DeleteMessage(id, uid); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
