package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
)

type LoginResponse struct {
	User  autocode.MinappUser `json:"userinfo"`
	Token string              `json:"token"`
}
