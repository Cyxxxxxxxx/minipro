package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"minipro/models"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CodeStruct struct {
	Code string `json:"code" form:"code"`
}

type CheckTokenResp struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	Userid      string `json:"Userid"`
	DeviceId    string `json:"DeviceId"`
	user_ticket string `json:"user_ticket"`
}

func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

type RespUserID struct {
	Code   string `json:"code"`
	Userid string `json:"userid"`
}

// HomeHandle 企业微信页面跳转接口
// @Summary 获取扫码应用请求地址
// @Description 使用本接口返回的地址作为该应用的请求地址！！
// @Tags 获取应用地址
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsgList
// @Router /home [get]
func HomeHandle(c *gin.Context) {
	url := UrlEncode("http://fanxiao.gzhu.edu.cn:8080/redirect_uri")
	login_url := fmt.Sprintf(`https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect`,
		"wxeecd4f1f44f855da",
		url,
	)
	ResponseSuccess(c, CodeSuccess, login_url)
}

// SetUseridHandle 获取code之后根据code解析调用人id 再插入上下文
func SetUseridHandle(c *gin.Context) {
	// 获取query中的code
	p := &CodeStruct{}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("c.ShouldBindQuery(p) failed ! 无法向企业微信获取Code！")
		c.JSON(http.StatusOK, "无法向企业微信获取Code！")
		return
	}
	zap.L().Info("CODE = ", zap.String("p.Code = ", p.Code))
	token, err := getAccessToken(1)
	if err != nil {
		zap.L().Error(" getAccessToken() failed err = ", zap.Error(err))
		c.JSON(http.StatusOK, "token无法获取！！！")
		return
	}

	userid, err := getMemberMsg(token, p.Code)
	if err != nil {
		zap.L().Error("getMemberMsg(token,p.Code) failed err = ", zap.Error(err))
		c.JSON(http.StatusOK, "无法获得userid！！")
		return
	}

	zap.L().Info("userid = ", zap.String("userid = ", userid))

	ResponseSuccess(c, CodeSuccess, userid)
}

// getAccessToken 换取token(扫码应用的)
func getAccessToken(a_type int) (string, error) {
	url := `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s`
	if a_type == 1 {
		// 合成url
		url = fmt.Sprintf(url, "wxeecd4f1f44f855da", "MApU9wrJKk_dB0NK8DqGVeW0sMThbvr84xeXBEoP8eE")
	} else if a_type == 2 {
		url = fmt.Sprintf(url, "wxeecd4f1f44f855da", "MT4_TyoU7iiA5GExU_T-Tm6DTvvyeQYIKOA_ftVMrUE")
	}
	// 创建http请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析http请求中的body数据 绑定到我们定义的结构体中
	wxResp := models.AccessToken{}
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&wxResp); err != nil {
		return "", err
	}
	// 判断是否 返回一个异常情况
	if wxResp.ErrCode != 0 {
		// 非0 表示出现异常情况
		return "", errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}
	return wxResp.Token, nil
}

// getMemberMsg 根据token和code 获取访问者信息
func getMemberMsg(token, code string) (string, error) {
	url := `https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s`
	// 合成url
	url = fmt.Sprintf(url, token, code)
	// 创建http请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析http请求中的body数据 绑定到我们定义的结构体中
	wxResp := CheckTokenResp{}
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&wxResp); err != nil {
		return "", err
	}
	// 判断是否 返回一个异常情况
	if wxResp.ErrCode != 0 {
		// 非0 表示出现异常情况
		return "", errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}

	msg, _ := json.Marshal(&wxResp)
	fmt.Println("menber message = ", string(msg))

	return wxResp.Userid, nil
}

// HomeHandleWeb 企业微信管理后台页面跳转接口
// @Summary 获取管理后台应用请求地址
// @Description 使用本接口返回的地址作为该应用的请求地址！！
// @Tags 获取应用地址
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsgList
// @Router /web [get]
func HomeHandleWeb(c *gin.Context) {
	url := UrlEncode("http://fanxiao.gzhu.edu.cn:8080/webRedirect_uri")
	login_url := fmt.Sprintf(`https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect`,
		"wxeecd4f1f44f855da",
		url,
	)
	ResponseSuccess(c, CodeSuccess, login_url)
}

// SetUseridHandleWeb 管理后台获得用户信息
func SetUseridHandleWeb(c *gin.Context) {
	//c.JSON(http.StatusOK, "我跳到了SetUseridHandle！！！！")
	// 获取query中的code
	p := &CodeStruct{}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("c.ShouldBindQuery(p) failed ! 无法向企业微信获取Code！")
		c.JSON(http.StatusOK, "无法向企业微信获取Code！")
		return
	}
	zap.L().Info("CODE = ", zap.String("p.Code = ", p.Code))
	token, err := getAccessToken(2)
	if err != nil {
		zap.L().Error(" getAccessToken() failed err = ", zap.Error(err))
		c.JSON(http.StatusOK, "token无法获取！！！")
		return
	}

	userid, err := getMemberMsg(token, p.Code)
	if err != nil {
		zap.L().Error("getMemberMsg(token,p.Code) failed err = ", zap.Error(err))
		c.JSON(http.StatusOK, "无法获得userid！！")
		return
	}

	zap.L().Info("userid = ", zap.String("userid = ", userid))
	ResponseSuccess(c, CodeSuccess, userid)

}
