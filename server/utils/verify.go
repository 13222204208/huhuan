package utils

var (
	IdVerify               = Rules{"ID": {NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify         = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}, "ParentId": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}

	//微信小程序登录验证
	MinAppLoginVerify = Rules{"Code": {NotEmpty()}, "Nickname": {NotEmpty()}, "Avatar": {NotEmpty()}}

	//微信小程序轮播图
	MinAppBannerVerify = Rules{"CityId": {NotEmpty()}, "ClassId": {NotEmpty()}}

	//创建小程序地址
	MinAppAddressVerify = Rules{"Country": {NotEmpty()}, "Province": {NotEmpty()}, "City": {NotEmpty()}, "Detail": {NotEmpty()}, "PostCode": {Eq("6")}}

	//投诉订单信息
	MinAppComplaintOrder = Rules{"Label": {NotEmpty()}, "Contents": {NotEmpty()}, "OrderNum": {NotEmpty()}, "ByComplainantId": {NotEmpty()}}

	//提交帮我修
	MinAppHelpRepair = Rules{"Contents": {NotEmpty()}, "Linkman": {NotEmpty()}, "ContactPhone": {RegexpMatch("phone")}, "Address": {NotEmpty()}, "Time": {NotEmpty()}, "Price": {NotEmpty()}}

	//提交帮我做订单
	MinAppHelpWork = Rules{"Contents": {NotEmpty()}, "Linkman": {NotEmpty()}, "ContactPhone": {NotEmpty()}, "Address": {NotEmpty()}, "Time": {NotEmpty()}, "Price": {NotEmpty()}}

	//提交帮我取订单
	MinAppHelpFetch = Rules{"GoodsName": {NotEmpty()}, "GoodsCount": {NotEmpty()}, "Linkman": {NotEmpty()}, "ContactPhone": {RegexpMatch("phone")}, "FetchAddress": {NotEmpty()}, "ReceiveAddress": {NotEmpty()}, "Time": {NotEmpty()}, "Price": {NotEmpty()}}

	//提交帮我买
	MinAppHelpBuy = Rules{"GoodsName": {NotEmpty()}, "GoodsCount": {NotEmpty()}, "Linkman": {NotEmpty()}, "ContactPhone": {RegexpMatch("phone")}, "BuyAddress": {NotEmpty()}, "ReceiveAddress": {NotEmpty()}, "Time": {NotEmpty()}, "Price": {NotEmpty()}, "GoodsPrice": {NotEmpty()}}

	//创建国内捎带 寻找乘机人
	MinAppHelpIncidentallyTypeOne = Rules{"Type": {NotEmpty()}, "GoodsName": {NotEmpty()}, "Price": {NotEmpty()}, "Linkman": {NotEmpty()}, "ContactPhone": {RegexpMatch("phone")}, "FetchAddress": {NotEmpty()}, "ReceiveAddress": {NotEmpty()}, "Time": {NotEmpty()}, "FlightNumber": {NotEmpty()}}
	//创建国内捎带 我是乘机人
	MinAppHelpIncidentallyTypeTwo = Rules{"Type": {NotEmpty()}, "Price": {NotEmpty()}, "Linkman": {NotEmpty()}, "ContactPhone": {RegexpMatch("phone")}, "StartCity": {NotEmpty()}, "ArriveCity": {NotEmpty()}, "Time": {NotEmpty()}, "FlightNumber": {NotEmpty()}}

	//组团拼车，我是车主
	MinAppHelpCarpoolTypeOne = Rules{"Type": {NotEmpty()}, "OwnerName": {NotEmpty()}, "Gender": {NotEmpty()}, "ContactPhone": {RegexpMatch("phone")}, "CarNumber": {NotEmpty()}, "UseNum": {NotEmpty()}, "StartAddress": {NotEmpty()}, "ArriveAddress": {NotEmpty()}, "CarType": {NotEmpty()}, "Time": {NotEmpty()}, "Price": {NotEmpty()}}

	//组团拼车，我是乘车人
	MinAppHelpCarpoolTypeTwo = Rules{"RiderName": {NotEmpty()}, "Gender": {NotEmpty()}, "ContactPhone": {RegexpMatch("phone")}, "UseNum": {NotEmpty()}, "UseTime": {NotEmpty()}, "StartAddress": {NotEmpty()}, "ArriveAddress": {NotEmpty()}, "CarType": {NotEmpty()}, "Price": {NotEmpty()}}

	//团购接龙发布商品
	MinAppHelpGroupon = Rules{"GoodsName": {NotEmpty()}, "GoodsTotal": {NotEmpty()}, "Price": {NotEmpty()}, "StartTime": {NotEmpty()}, "OverTitle": {NotEmpty()}, "GoodsImg": {NotEmpty()}}

	//二手闲置商品发布
	MinAppHelpSecond = Rules{"GoodsDetail": {NotEmpty()}, "GoodsImg": {NotEmpty()}, "GoodsName": {NotEmpty()}, "OriginalPrice": {NotEmpty()}, "Price": {NotEmpty()}, "Address": {NotEmpty()}, "Category": {NotEmpty()}}

	//司机认证
	CheckDriverVerify = Rules{"Name": {NotEmpty()}, "Phone": {RegexpMatch("phone")}, "CarType": {NotEmpty()}, "Gender": {NotEmpty()}, "CarImage": {NotEmpty()}, "Licence": {NotEmpty()}}

	//商家认证
	CheckMerchantVerify = Rules{"Name": {NotEmpty()}, "Phone": {RegexpMatch("phone")}, "Licence": {NotEmpty()}, "Charter": {NotEmpty()}}

	//二手闲置留言和评论
	SecondMessageVerify = Rules{"OrderNum": {NotEmpty()}, "Tid": {NotEmpty()}, "Contents": {NotEmpty()}}

	//微信用户身份验证
	CheckIdentifyVerify = Rules{"Name": {NotEmpty()}, "Phone": {RegexpMatch("phone")}, "gender": {NotEmpty()}, "Card": {NotEmpty()}, "Selfie": {NotEmpty()}}
)
