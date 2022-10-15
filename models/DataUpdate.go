package models

import (
	"fmt"
	"time"
)

const (
	StatusNotInSch   = "1"  //
	StatusInSch      = "2"  // 已返校
	StatusLateToSch  = "3"  // 逾期返校
	StatusDelayToSch = "4"  // 延缓返校
	StatusNotApply   = "-1" // 未申请返校
	FlashNotInSch    = "20"
	FlashInSch       = "21"
	FlashLateToSch   = "22"
)

var DontInsertAgain = fmt.Errorf("请勿重复插入相同学号的学生")

type StudentMsg struct {
	Name        string `json:"name" db:"O_STUDENT_BASIC_NAME"`                // 学生姓名
	StudentNum  int64  `json:"student_num" db:"O_STUDENT_BASIC_STUDENTID"`    // 学生学号
	StudentCps  string `json:"student_cps" db:"O_STUDENT_BASIC_CAMPUS"`       // 学生所在校区
	CollegeName string `json:"college_name" db:"O_STUDENT_BASIC_COLLEGENAME"` // 学院名称
	Grade       string `json:"grade" db:"O_STUDENT_BASIC_GRADE"`              // 所在年级
}

type StuStatusResp struct {
	ArrivalStatus    string    `json:"arrival_status" db:"arrival_status"`
	ExpectedBackTime time.Time `json:"expected_back_time" db:"expected_back_time"`
}

type CreateStudent struct {
	StudentId         string `json:"student_id" db:"student_id" binding:"required" example:"2006100141"`                  // 学生学号
	ExceptArrivalTime string `json:"except_arrival_time" db:"expected_back_time" binding:"required" example:"2022-09-01"` // 预计到达时间 “yyyy-mm-dd”传入
}
