package request

type SecondRoomMessageRequest struct {
	RoomId uint `json:"roomId" form:"roomId"`
}
