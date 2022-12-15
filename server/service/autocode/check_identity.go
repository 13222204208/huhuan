package autocode

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
)

type CheckIdentityService struct {
}

// CreateCheckIdentity 创建CheckIdentity记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkIdentityService *CheckIdentityService) CreateCheckIdentity(checkIdentity autocode.CheckIdentity) (err error) {
	if !errors.Is(global.GVA_DB.Where("uid", checkIdentity.Uid).First(&checkIdentity).Error, gorm.ErrRecordNotFound) {
		return errors.New("已经提交了身份验证")
	}
	fmt.Println("身份证号", checkIdentity.Card, "姓名", checkIdentity.Name)
	info := GetUrl(checkIdentity.Name, checkIdentity.Card)
	if info.Success == false {
		return errors.New(info.Msg)
	}

	err = global.GVA_DB.Create(&checkIdentity).Error

	//修改用户是否提交了申请
	var user autocode.MinappUser
	global.GVA_DB.Model(&user).Where("id", checkIdentity.Uid).UpdateColumn("verify", 1)

	return err
}

type result struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
}

func GetUrl(name, card string) (res *result) {
	urlLink := "https://idcardv2.shumaidata.com/new-idcard?idcard=" + card + "&name=" + url.QueryEscape(name)

	appcode := "1a1223e2e868408e92d2ca370bb4cc5e"
	header := "APPCODE " + appcode

	client := http.Client{}
	r, err := http.NewRequest(http.MethodGet, urlLink, nil)
	if err != nil {
		fmt.Println(err, 1)
	}
	// 添加请求头

	r.Header.Add("Authorization", header)
	fmt.Println("请求的url", urlLink)
	// 发送请求
	resp, err := client.Do(r)
	fmt.Println("发送的数据", resp)
	if err != nil {
		fmt.Println(err, 2)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err, 3)
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		panic(err)
	}
	fmt.Println("验证身份证号", string(b))
	return res
}

// DeleteCheckIdentity 删除CheckIdentity记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkIdentityService *CheckIdentityService) DeleteCheckIdentity(checkIdentity autocode.CheckIdentity) (err error) {
	err = global.GVA_DB.Delete(&checkIdentity).Error
	return err
}

// DeleteCheckIdentityByIds 批量删除CheckIdentity记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkIdentityService *CheckIdentityService) DeleteCheckIdentityByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CheckIdentity{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCheckIdentity 更新CheckIdentity记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkIdentityService *CheckIdentityService) UpdateCheckIdentity(checkIdentity autocode.CheckIdentity) (err error) {
	err = global.GVA_DB.Save(&checkIdentity).Error
	//如果审核的状态为 1 代表未审核，2 代表审核通过，3 审核拒绝
	if err != nil {
		return err
	} else {
		var checkState int
		if checkIdentity.Status == 2 {
			checkState = 1
		}
		var user autocode.MinappUser

		err = global.GVA_DB.Model(&user).Where("id", checkIdentity.Uid).Updates(map[string]interface{}{"name": checkIdentity.Name, "phone": checkIdentity.Phone, "verify": checkState}).Error
		return err
	}
}

// GetCheckIdentity 根据id获取CheckIdentity记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkIdentityService *CheckIdentityService) GetCheckIdentity(id uint) (err error, checkIdentity autocode.CheckIdentity) {
	err = global.GVA_DB.Where("id = ?", id).First(&checkIdentity).Error
	return
}

// GetCheckIdentityInfoList 分页获取CheckIdentity记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkIdentityService *CheckIdentityService) GetCheckIdentityInfoList(info autoCodeReq.CheckIdentitySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.CheckIdentity{})
	var checkIdentitys []autocode.CheckIdentity
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("MinAppUser").Find(&checkIdentitys).Error
	return err, checkIdentitys, total
}
