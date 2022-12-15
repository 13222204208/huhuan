package request

type BannerList struct {
	City  string `json:"city" form:"city"`
	Class int    `json:"class" form:"class"`
}
