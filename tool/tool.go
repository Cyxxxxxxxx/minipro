package tool

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"minipro/models"
	"net/http"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetTokenModule struct {
	Clientid string `json:"clientid"`
	Secretid string `json:"secretid"`
}

type GetResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	//ResultData `json:"resultData"`
}

type ResultData struct {
	Token string `json:"token"`
}

type UpdataMsgModule struct {
	Idserial  string `json:"idserial"`
	Typecode  string `json:"typecode"`
	Statecode string `json:"statecode"`
}

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

var ErrorUserNotExist = errors.New("扫码人id错误")

// SetWildCardBtwWord 在字符串前后加上  %
func SetWildCardBtwWord(word string) string {
	c := "%"
	return c + word + c
}

// SetWildCardInSta 处理绑定的字符串
func SetWildCardInSta(p *models.MsgList) *models.MsgList {
	p.StudentName = SetWildCardBtwWord(p.StudentName)
	p.StudentNum = SetWildCardBtwWord(p.StudentNum)
	p.Campus = SetWildCardBtwWord(p.Campus)
	p.CollegeName = SetWildCardBtwWord(p.CollegeName)
	p.Grade = SetWildCardBtwWord(p.Grade)
	p.ArrivalTime = SetWildCardBtwWord(p.ArrivalTime)
	return p
}

//接收传来的body
func GetJsonMap(c *gin.Context) (map[string]interface{}, error) {
	buf, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		return nil, err
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)

	return jsonMap, err
}

// UpdateBTSstatus 更新返校状态码 22 -- 已返校  21 -- 逾期
func UpdateBTSstatus(userid string, code string, iscreate bool) (bool, error) {
	var url string
	if iscreate {
		url = `http://wx_demo:8000/openPlatform/callSaveSysQrcode`
	} else {
		url = `http://wx_demo:8000/openPlatform/callModifySysQrcode`
	}
	msg := UpdataMsgModule{
		Idserial:  userid,
		Typecode:  "2",
		Statecode: code,
	}

	bytesData, err := json.MarshalIndent(&msg, "", "\t\t")
	if err != nil {
		zap.L().Error("json.MarshalIndent failed = ", zap.Error(err))
		return false, err
	}

	// 创建http请求
	resp, _ := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(bytesData))

	defer func() {
		if e := recover(); e != nil {
			zap.L().Error("!!!! err = recover() 发送post请求更新或创建校园码失败", zap.Error(nil))
			return
		}
		resp.Body.Close()
	}()

	// resp 可用 ; 解析http请求中的body数据 绑定到我们定义的结构体中
	Resp := GetResp{}
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&Resp); err != nil {
		zap.L().Error("decoder.Decode(&wxResp) failed err : ", zap.Error(err))
		return false, err
	}
	if !Resp.Success {
		zap.L().Error("Resp.Success = false message =  ", zap.Error(fmt.Errorf("%s", Resp.Message)))
		return false, err
	}
	return true, nil
}

// SturctToByteSlice 结构体转为二进制
func SturctToByteSlice(msg *UpdataMsgModule) []byte {
	Len := unsafe.Sizeof(*msg)
	testBytes := &SliceMock{
		addr: uintptr(unsafe.Pointer(msg)),
		cap:  int(Len),
		len:  int(Len),
	}
	return *(*[]byte)(unsafe.Pointer(testBytes))
}

// InitInMain 初始化时区
func InitInMain() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}
