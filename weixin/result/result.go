package result

//CommonResult 微信返回的公共结果对象
type CommonResult struct {
	Errcode int    //返回code
	Errmsg  string //返回的描述
}

//TokenResult 微信全局token返回的结果对象
type TokenResult struct {
	CommonResult
	AccessToken string //token信息
	ExpiresIn   int64  //token的过期时间

}

//TicketResult 微信返回的票据信息对象
type TicketResult struct {
	CommonResult
	Ticket    string //票据信息
	ExpiresIn int64  //票据过期时间
}

//ShortResult 微信返回的短链接信息
type ShortResult struct {
	CommonResult
	ShortUrl string //短链接信息
}

//UserToken 用户授权的token信息
type UserToken struct {
	CommonResult
	AccessToken   string //授权token
	ExpiresIn     string //过期时间
	ResfreshToken string //刷新token的凭证
	OpenId        string //用户的openId
	Scope         string //作用域
	UnionId       string //公开，基于公众账号的全局唯一id
}

// UserResult 微信用户信息
type UserResult struct {
	CommonResult
	Id         int64  //微信用户主键Id
	Subscribe  int    //用户是否订阅该公众账号，0：未关注
	OpenId     string //用户的标识，当前公众账号唯一
	NickName   string //用户的昵称
	Sex        int    //用户的性别，1男，2女，0未知
	Lauguage   string //语言
	City       string //城市
	Province   string //省份
	Country    string //国家
	HeadImgUrl string //头像地址
	/*
	 * 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。 详见：获取用户个人信息 （ UnionID机制 {@link
	 * https://open
	 * .weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource
	 * /res_list&verify=1&lang=zh_CN}）
	 */
	UnionId string
	Remark  string //公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupId string //用户所在分组
}
