package controller

import (
	"minipro/models"
	"minipro/mysql"
	"minipro/service"
)

type _ResponseMsgList struct {
	Code    ResCode                  `json:"code"`    // 业务响应状态码
	Message string                   `json:"message"` // 提示信息
	Data    []*models.ShowStudentMsg `json:"data"`    // 数据
}

type _ResponseDataShow struct {
	Code    ResCode              `json:"code"`    // 业务响应状态码
	Message string               `json:"message"` // 提示信息
	Data    []*models.StudentMsg `json:"data"`    // 数据
}

type _ResponseData struct {
	Code    ResCode     `json:"code"`    // 业务响应状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 数据
}

type _ResponseUpdateDate struct {
	Code    ResCode `json:"code"`    // 业务响应状态码
	Message string  `json:"message"` // 提示信息
}

//
type RespSuccess struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type RespStaInfo struct {
	Code int64            `json:"code"`
	Msg  string           `json:"msg"`
	Data service.NumOfStu `json:"data"`
}

type RespStaffInfo struct {
	Code int64                 `json:"code"`
	Msg  string                `json:"msg"`
	Data mysql.StaffInfoStruct `json:"data"`
}

type RespStaff struct {
	Code int64               `json:"code"`
	Msg  string              `json:"msg"`
	Data mysql.RespStaffInfo `json:"data"`
}

type RespCollege struct {
	Code int64    `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

type RespAuth struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data int64  `json:"data"`
}

type _ResponseSearch struct {
	Code    ResCode       `json:"code"`    // 业务响应状态码
	Message string        `json:"message"` // 提示信息
	Data    []*models.Stu `json:"data"`    // 数据
}
type _ResponseGetStaySchoolList struct {
	Code    ResCode       `json:"code"`    // 业务响应状态码
	Message string        `json:"message"` // 提示信息
	Data    []interface{} `json:"data"`    // 数据
}
type _ResponseSucc struct {
	Code    ResCode          `json:"code"`    // 业务响应状态码
	Message string           `json:"message"` // 提示信息
	Data    []*ResponeseCode `json:"data"`    // 数据
}

//简单响应只返回code和message
type _ResponseMsg struct {
	Code    ResCode `json:"code"`    // 业务响应状态码
	Message string  `json:"message"` // 提示信息
}
type _ResponseList struct {
	Code       ResCode       `json:"code"`    // 业务响应状态码
	Message    string        `json:"message"` // 提示信息
	Data       []*models.Stu `json:"data"`    // 数据
	TotalQuery int           `json:"totalQuery"`
}
