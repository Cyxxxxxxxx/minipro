package models

import "time"

type MsgList struct {
	Page        int64  `form:"page" example:"1"`                 // 页号
	Size        int64  `form:"size" example:"10"`                // 条数
	Type        int64  `form:"type" example:"1"`                 //
	StudentNum  string `form:"studentnum" example:"2006100141"`  // 学号
	StudentName string `form:"studentname" example:"姜"`          // 姓名
	Campus      string `form:"campus" example:"大学城"`             // 校区
	CollegeName string `form:"collegename" example:"计算机"`        // 学院
	Grade       string `form:"grade" example:"2020"`             // 年级
	ArrivalTime string `form:"arrivaltime" example:"2020-09-01"` // 计划返校日期
}

type StudentMsgUg struct {
	Name        string `json:"name" db:"O_STUDENT_BASIC_NAME"`                // 学生姓名
	StudentNum  int64  `json:"student_num" db:"student_id"`                   // 学生学号
	StudentCps  string `json:"student_cps" db:"O_STUDENT_BASIC_CAMPUS"`       // 学生所在校区
	CollegeName string `json:"college_name" db:"O_STUDENT_BASIC_COLLEGENAME"` // 学院名称
	Grade       string `json:"grade" db:"O_STUDENT_BASIC_GRADE"`              // 学生年级
	StuStatus   int    `json:"stu_status" db:"arrival_status"`                // 返校状态
	ArrivalTime string `json:"arrival_time" db:"expected_back_time"`          // 预计返校时间
	ScanId      string `json:"scan_id" db:"scan_id"`                          // 扫码者id
}

type StudentMsgUgt struct {
	Name        string `json:"name" db:"O_STUDENT_BASIC_NAME"`                // 学生姓名
	StudentNum  int64  `json:"student_num" db:"student_id"`                   // 学生学号
	StudentCps  string `json:"student_cps" db:"O_STUDENT_BASIC_CAMPUS"`       // 学生所在校区
	CollegeName string `json:"college_name" db:"O_STUDENT_BASIC_COLLEGENAME"` // 学院名称
	Grade       string `json:"grade" db:"O_STUDENT_BASIC_GRADE"`              // 学生年级
	StuStatus   int    `json:"stu_status" db:"arrival_status"`                // 返校状态
	ArrivalTime string `json:"arrival_time" db:"arrival_time"`                // 到达时间
	ScanId      string `json:"scan_id" db:"scan_id"`                          // 扫码者id
}

type StudentMsgUgw struct {
	Name        string `json:"name" db:"O_STUDENT_BASIC_NAME"`                // 学生姓名
	StudentNum  int64  `json:"student_num" db:"O_STUDENT_BASIC_STUDENTID"`    // 学生学号
	StudentCps  string `json:"student_cps" db:"O_STUDENT_BASIC_CAMPUS"`       // 学生所在校区
	CollegeName string `json:"college_name" db:"O_STUDENT_BASIC_COLLEGENAME"` // 学院名称
	Grade       string `json:"grade" db:"O_STUDENT_BASIC_GRADE"`              // 学生年级
	StuStatus   int    `json:"stu_status" db:"arrival_status"`                // 返校状态
	ArrivalTime string `json:"arrival_time" db:"expected_back_time"`          // 预计返校时间
	ScanId      string `json:"scan_id" db:"scan_id"`                          // 扫码者id
}

type StudentMsgPg struct {
	Name        string `json:"name" db:"O_STUDENT_BASIC_NAME"`                 // 学生姓名
	StudentNum  int64  `json:"student_num" db:"O_STUDENT_BASIC_STUDENTID"`     // 学生学号
	StudentCps  string `json:"student_cps" db:"O_STUDENT_BASIC_XQMC"`          // 学生所在校区
	CollegeName string `json:"college_name" db:"O_STUDENT_BASIC_COLLEGENAME"`  // 学院名称
	Grade       string `json:"grade" db:"O_STUDENT_BASIC_ENROLLYEAR"`          // 学生年级
	StuStatus   int    `json:"stu_status,omitempty" db:"arrival_status"`       // 返校状态
	ArrivalTime string `json:"arrival_time,omitempty" db:"expected_back_time"` // 预计返校时间
	ScanId      string `json:"scan_id,omitempty" db:"scan_id"`                 // 扫码者id
}

type StudentMsgPgt struct {
	Name        string `json:"name" db:"O_STUDENT_BASIC_NAME"`                // 学生姓名
	StudentNum  int64  `json:"student_num" db:"O_STUDENT_BASIC_STUDENTID"`    // 学生学号
	StudentCps  string `json:"student_cps" db:"O_STUDENT_BASIC_XQMC"`         // 学生所在校区
	CollegeName string `json:"college_name" db:"O_STUDENT_BASIC_COLLEGENAME"` // 学院名称
	Grade       string `json:"grade" db:"O_STUDENT_BASIC_ENROLLYEAR"`         // 学生年级
	StuStatus   int    `json:"stu_status,omitempty" db:"arrival_status"`      // 返校状态
	ArrivalTime string `json:"arrival_time,omitempty" db:"arrival_time"`      // 预计返校时间
	ScanId      string `json:"scan_id,omitempty" db:"scan_id"`                // 扫码者id

}

type StudentMsgPgw struct {
	Name        string `json:"name" db:"O_STUDENT_BASIC_NAME"`                // 学生姓名
	StudentNum  int64  `json:"student_num" db:"O_STUDENT_BASIC_STUDENTID"`    // 学生学号
	StudentCps  string `json:"student_cps" db:"O_STUDENT_BASIC_XQMC"`         // 学生所在校区
	CollegeName string `json:"college_name" db:"O_STUDENT_BASIC_COLLEGENAME"` // 学院名称
	Grade       string `json:"grade" db:"O_STUDENT_BASIC_ENROLLYEAR"`         // 学生年级
	StuStatus   int    `json:"stu_status" db:"arrival_status"`                // 返校状态
	ArrivalTime string `json:"arrival_time" db:"expected_back_time"`          // 预计返校时间
	ScanId      string `json:"scan_id" db:"scan_id"`                          // 扫码者id

}

type ShowStudentMsg struct {
	Name        string `json:"name"`                   // 学生姓名
	StudentNum  int64  `json:"student_num"`            // 学生学号
	StudentCps  string `json:"student_cps"`            // 学生所在校区
	CollegeName string `json:"college_name"`           // 学院名称
	Grade       string `json:"grade"`                  // 学生年级
	StuStatus   int    `json:"stu_status,omitempty"`   // 返校状态
	ArrivalTime string `json:"arrival_time,omitempty"` // 预计返校时间
	ScanId      string `json:"scan_id,omitempty"`      // 确认者id
}

type StudStatus struct {
	Status     int       `json:"status" db:"arrival_status"`
	ExceptTime time.Time `json:"except_time" db:"expected_back_time"`
}
