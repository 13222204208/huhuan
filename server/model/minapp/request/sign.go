package request

type ChangeSignInState struct {
	State string `json:"state"`
}

type ExchangeCoupon struct {
	CouponId uint `json:"couponId"`
}
