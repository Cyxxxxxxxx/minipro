package controller

//定义一些可能出现的错误码

type ResCode int64

const (
	CodeSuccess ResCode = 500 + iota
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:    "success",
	CodeServerBusy: "错误",
}

// Msg 返回特定的错误提示信息
func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
