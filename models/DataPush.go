package models

import "time"

// AccessToken 请求接口返回带有token的JSON数据
type AccessToken struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Token     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
}

// Data 发送给企业微信的JSON数据
type Data struct {
	Touser                 string `json:"touser"`
	MsgType                string `json:"msgtype"`
	AgentId                int    `json:"agentid"`
	Text                   Text   `json:"text"`
	Safe                   int    `json:"safe"`
	EnableIDTrans          int    `json:"enable_id_trans"`
	EnableDuplicateCheck   int    `json:"enable_duplicate_check"`
	DuplicateCheckInterval int    `json:"duplicate_check_interval"`
}

// Text 发送数据的文本
type Text struct {
	Content string `json:"content"`
}

// PushMessageResp 消息推送返回的结构体
type PushMessageResp struct {
	ErrCode        int    `json:"errcode"`
	ErrMsg         string `json:"errmsg"`
	Invaliduser    string `json:"invaliduser"`
	Unlicenseduser string `json:"unlicenseduser"`
	MsgID          string `json:"msgid"`
	ResponseCode   string `json:"response_code"`
}

// StudentBackMsg 学生返校信息
type StudentBackMsg struct {
	StudentID        string    `json:"student_id" db:"student_id"`                 // 学生学号
	StudentName      string    `json:"student_name" db:"O_STUDENT_BASIC_NAME"`     // 学生姓名
	ExpectedBackTime time.Time `json:"expected_back_time" db:"expected_back_time"` // 学生预计返校时间

}
