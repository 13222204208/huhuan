package request

type SendSystemMessageRequest struct {
	Uid      []string `json:"uid" form:"uid"`
	Resource string   `json:"resource" form:"resource"`
	Contents string   `json:"contents" form:"contents"`
}
