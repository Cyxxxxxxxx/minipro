package models

//查询对应的结构体
type Stu struct {
	O_STUDENT_BASIC_NAME        string `json:"name"`        //姓名
	O_STUDENT_BASIC_STUDENTID   string `json:"stuid"`       //学号
	O_STUDENT_BASIC_COLLEGENAME string `json:"collegename"` //学院
	O_STUDENT_CATEGORY          string `json:"category"`    //学生类别 (本科生或研究生)
	O_STUDENT_BASIC_GRADE       string `json:"grade"`       //年级
	O_STUDENT_BASIC_CAMPUS      string `json:"campus"`      //校区(本科生表)
	Submit_time                 string `json:"-"`           //申请提交时间----用来判断学生状态（留校，未留校）
	O_STUDENT_BASIC_XQMC        string `json:"-"`           //校区(研究生表)
	O_STUDENT_STATUS            string `json:"status"`      //学生状态 (留校或未留校)
	O_STUDENT_BASIC_EDU_LEVEL   string `json:"-"`           //区分硕士与博士
}
type Stu_back struct {
	Stu_id string `json:"stu_id"` //
}

//生成留校学生表字符串  例如：stay_school_stu_2022_2023_spring/autumn
func SetTableName(y, s string) string {
	var S string
	if s == "秋季学期" {
		S = "autumn"
	} else {
		S = "spring"
	}
	return "stay_school_stu_" + y + "_" + S
}

//返回导入成功和失败人数的结构体
type Res01 struct {
	Code   int    `json:"Code"`
	Show01 string `json:"Show"`
}

//对应未返校数据
type NoReturnedData struct {
	Name        string `json:"name" binding:"required"`
	Number      int64  `json:"student_num" binding:"required"`
	CollegeName string `json:"college_name" binding:"required"`
	Grade       string `json:"grade" binding:"required"`
	Campus      string `json:"student_cps" binding:"required"`
	Status      int64  `json:"stu_status" binding:"required"`
	PlanDate    string `json:"arrival_time" binding:"required"`
}

//对应已返校数据
type ReturnedData struct {
	Name          string `json:"name" binding:"required"`
	Number        int64  `json:"student_num" binding:"required"`
	CollegeName   string `json:"college_name" binding:"required"`
	Grade         string `json:"grade" binding:"required"`
	Campus        string `json:"student_cps" binding:"required"`
	Status        int64  `json:"stu_status" binding:"required"`
	PlanDate      string `json:"arrival_time" binding:"required"`
	ConfirmPerson string `json:"scan_id" binding:"required"`
}

//对应未申请数据
type NoFiledData struct {
	Name        string `json:"name" binding:"required"`
	Number      int64  `json:"student_num" binding:"required"`
	CollegeName string `json:"college_name" binding:"required"`
	Grade       string `json:"grade" binding:"required"`
	Status      int64  `json:"stu_status" binding:"required"`
}

type Inter struct {
	Face []interface{}
}
