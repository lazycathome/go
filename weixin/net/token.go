package net

import (
	"encoding/json"
	"lazycathome/weixin/cache"
	"lazycathome/weixin/conf"
	"lazycathome/weixin/result"

	"strconv"

	"fmt"

	"github.com/go-resty/resty"
)

const (
	accessTokenURL string = "https://api.weixin.qq.com/cgi-bin/token"
	ticketURL      string = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi"
	expireTime     int    = 7100
)

//GetToken 获取token信息
func GetToken(wxConf conf.WXConf) (string, error) {
	param := make(map[string]interface{})
	param["appId"] = wxConf.AppId
	param["appSecret"] = wxConf.AppSecret
	body, err := post(accessTokenURL, param)
	if err != nil {
		return "", err
	}
	var tr result.TokenResult
	m := make(map[string]string)
	json.Unmarshal(body, &m)
	errCode, err := strconv.Atoi(m["errcode"])
	if err != nil {
		errCode = -1
		fmt.Println(err)
	}
	tr.Errcode = errCode
	ticket := m["accesstoken"]
	if errCode == 0 && ticket != "" {
		cache.SetGobalToken("gobalToken", ticket, expireTime)
	}

	return ticket, err
}

//GetTicket 获取票据信息
func GetTicket(wxConf conf.WXConf) (string, error) {
	body, err := post(ticketURL, nil)
	if err != nil {
		return "", err
	}
	var tr result.TicketResult
	m := make(map[string]string)
	json.Unmarshal(body, &m)
	errCode, err := strconv.Atoi(m["errcode"])
	if err != nil {
		errCode = -1
		fmt.Println(err)
	}
	tr.Errcode = errCode
	ticket := m["ticket"]
	if errCode == 0 && ticket != "" {
		cache.SetGobalToken("gobalToken", ticket, expireTime)
	}

	return ticket, err
}

func post(url string, params map[string]interface{}) ([]byte, error) {
	resp, err := resty.R().
		SetBody(params).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func Get(url string, params map[string]string) ([]byte, error) {
	resp, err := resty.R().
		SetQueryParams(params).
		SetHeader("Accept", "application/json").
		Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
