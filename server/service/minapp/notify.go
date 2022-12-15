package minapp

import (
	"github.com/silenceper/wechat/v2/pay"
)

type NotifyService struct{}

func (s *NotifyService) PayNotify() {
	var u UserService
	u.PayConfig()
	var p pay.Pay
	p.GetNotify()
}
