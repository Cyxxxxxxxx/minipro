package models

import "image"

type HealthClockList struct {
	ID                       string `json:"id" db:"id"`                                                   //学号
	Gender                   string `json:"gander" db:"gander"`                                           //性别
	Phone                    string `json:"phone" db:"phone"`                                             //电话
	Monitor                  string `json:"monitor" db:"monitor"`                                         //班长
	Preceptor                string `json:"preceptor" db:"preceptor"`                                     //辅导员
	Origin                   string `json:"origin" db:"origin"`                                           //籍贯
	StayOrWorkSchool         string `json:"stay_or_work_school" db:"stay_or_work_school"`                 //校内居住/办公
	NowStayStatus            string `json:"now_stays_tatus" db:"now_stay_status"`                         //目前所在地状态(国内/国外)
	NowStay                  string `json:"now_stay" db:"now_stay"`                                       //目前所在地
	OutThisDay               string `json:"out_this_day" db:"out_this_day"`                               //当日是否外出
	HealthStatus             string `json:"health_status" db:"health_status"`                             //身体状态
	LastestAcid              string `json:"lastest_acid" db:"lastest_acid"`                               //最近一次核酸时间
	KeyPersonnel             string `json:"key_personnel" db:"key_personnel"`                             //本人是否为重点观察人员
	FamilyHealth             string `json:"family_health" db:"family_health"`                             //家人身体状态
	KeyFamily                string `json:"key_family" db:"key_family"`                                   //家人是否为重点观察人员
	DoubtfulOrDiagnosed      string `json:"doubtful_or_diagnosed" db:"doubtful_or_diagnosed"`             //疑似或确诊
	TouchDoubtfulOrDiagnosed string `json:"touch_doubtful_or_diagnosed" db:"touch_doubtful_or_diagnosed"` //是否接触过疑似或确诊
	HalfMonthTouch           string `json:"half_month_touch" db:"half_month_touch"`                       //是否接触过本个月内有疫情重点地区旅居史的人员
	GreenCode                string `json:"green_code" db:"green_code"`                                   //健康码是否为绿码
	HalfMonthGet             string `json:"half_month_get" db:"half_month_get"`                           //半个月内是否到过国内疫情重点地区
	Inoculation              string `json:"inoculation" db:"inoculation"`                                 //是否接种新冠疫苗
	InoculationNumber        string `json:"inoculation_number" db:"inoculation_number"`                   //接种的针次
	InoculationDate          string `json:"inoculation_date" db:"inoculation_date"`                       //疫苗接种日期
	InoculationMaker         string `json:"inoculation_maker" db:"inoculation_maker"`                     //疫苗提供厂商
	Pramise                  string `json:"pramise" db:"pramise"`                                         //个人承诺
	ClockDate                string `json:"clock_date" db:"clock_date"`                                   //本次打卡时间
}

//show HealthClockList
type ShowHealthClockList struct {
	Page                     int64  `json:"page" example:"1"`            //页号
	Size                     int64  `json:"size" example:"10"`           //条数
	Name                     string `json:"name"`                        //姓名
	ID                       string `json:"id"`                          //学号
	Grade                    string `json:"grade"`                       //年级
	Class                    string `json:"class"`                       //班级
	Gender                   string `json:"gender"`                      //性别
	Phone                    string `json:"phone"`                       //电话
	BronDate                 string `json:"bron_date"`                   //出生日期
	Academy                  string `json:"academy"`                     //学院
	Monitor                  string `json:"monitor"`                     //班长
	Preceptor                string `json:"preceptor"`                   //辅导员
	Origin                   string `json:"origin"`                      //籍贯
	StayOrWorkSchool         string `json:"stay_or_work_school"`         //校内居住/办公
	NowStayStatus            string `json:"now_stays_tatus"`             //目前所在地状态(国内/国外)
	NowStay                  string `json:"now_stay"`                    //目前所在地
	OutThisDay               string `json:"out_this_day"`                //当日是否外出
	HealthStatus             string `json:"health_status"`               //身体状态
	LastestAcid              string `json:"lastest_acid"`                //最近一次核酸时间
	KeyPersonnel             string `json:"key_personnel"`               //本人是否为重点观察人员
	FamilyHealth             string `json:"family_health"`               //家人身体状态
	KeyFamily                string `json:"key_family"`                  //家人是否为重点观察人员
	DoubtfulOrDiagnosed      string `json:"doubtful_or_diagnosed"`       //疑似或确诊
	TouchDoubtfulOrDiagnosed string `json:"touch_doubtful_or_diagnosed"` //是否接触过疑似或确诊
	HalfMonthTouch           string `json:"half_month_touch"`            //是否接触过本个月内有疫情重点地区旅居史的人员
	GreenCode                string `json:"green_code"`                  //健康码是否为绿码
	HalfMonthGet             string `json:"half_month_get"`              //半个月内是否到过国内疫情重点地区
	Inoculation              string `json:"inoculation"`                 //是否接种新冠疫苗
	InoculationNumber        string `json:"inoculation_number"`          //接种的针次
	InoculationDate          string `json:"inoculation_date"`            //疫苗接种日期
	InoculationMaker         string `json:"inoculation_maker"`           //疫苗提供厂商
	Pramise                  string `json:"pramise"`                     //个人承诺
	ClockDate                string `json:"clock_date"`                  //本次打卡时间
}

type ExcusedList struct {
	ID              string      `json:"id" db:"id"`                             //学号
	LeaveType       string      `json:"leave_type" db:"leave_type"`             //请假类型
	LeaveDays       string      `json:"leave_days" db:"leave_days"`             //请假天数
	LeaveTimeStart  string      `json:"leave_time_start" db:"leave_time_start"` //请假开始时间
	LeabeTimeEnd    string      `json:"leabe_time_end" db:"leave_time_end"`     //请假结束时间
	Photo           image.Image `json:"photo" db:"photo"`                       //照片附件
	LeaveSchool     string      `json:"leave_school" db:"leave_school"`         //是否离校
	OutWhere01      string      `json:"out_where01" db:"out_where01"`           //外出地点1
	OutWhere02      string      `json:"out_where02" db:"out_where02"`           //外出地点2
	OutWhere03      string      `json:"out_where03" db:"out_where03"`           //外出地点3
	Campus          string      `json:"campus" db:"campus:"`                    //校区
	PersonNumber    string      `json:"person_number" db:"person_number"`       //本人电话
	EmergencyNumber string      `json:"emergency_number" db:"emergency_number"` //紧急联系人电话
	LeaveReason     string      `json:"leave_reason" db:"leave_reason"`         //请假理由
	Status          string      `json:"status" db:"status"`                     //本次申请状体
	Preceptor       string      `json:"preceptor" db:"preceptor"`               //审核辅导员
}

//show ExcusedList
type ShowExcusedList struct {
	ID              string      `json:"id"`               //学号
	LeaveType       string      `json:"leave_type"`       //请假类型
	LeaveDays       string      `json:"leave_days"`       //请假天数
	LeaveTimeStart  string      `json:"leave_time_start"` //请假开始时间
	LeabeTimeEnd    string      `json:"leabe_time_end"`   //请假结束时间
	Photo           image.Image `json:"photo"`            //照片附件
	LeaveSchool     string      `json:"leave_school"`     //是否离校
	OutWhere01      string      `json:"out_where01"`      //外出地点1
	OutWhere02      string      `json:"out_where02"`      //外出地点2
	OutWhere03      string      `json:"out_where03"`      //外出地点3
	Campus          string      `json:"campus"`           //校区
	PersonNumber    string      `json:"person_number"`    //本人电话
	EmergencyNumber string      `json:"emergency_number"` //紧急联系人电话
	LeaveReason     string      `json:"leave_reason"`     //请假理由
	Status          string      `json:"status"`           //本次申请状体
	Preceptor       string      `json:"preceptor"`        //审核辅导员
}

type BackAppList struct {
	ID           string      `json:"id" db:"id"`                       //学号
	HealthStatus string      `json:"health_status" db:"health_status"` //健康状态
	FromWhere    string      `json:"from_where" db:"from_where"`       //从何地出发
	StartTime    string      `json:"start_time" db:"start_time"`       //出发时间
	GetTime      string      `json:"get_time" db:"get_time"`           //预计到校时间
	Vehicle      string      `json:"vehicle" db:"vehicle"`             //交通工具
	StaySchool   string      `json:"stay_school" db:"stay_school"`     //是否住校
	CarNumber    string      `json:"car_number" db:"car_number"`       //自驾车牌
	SuiKangMini  string      `json:"suikang_mini" db:"suikang_mini"`   //是否填报穗康小程序
	HealthPhoto  image.Image `json:"health_photo" db:"health_photo"`   //健康码
	TripPhoto    image.Image `json:"trip_photo" db:"trip_photo"`       //行程卡
	Preceptor    string      `json:"preceptor" db:"preceptor"`         //审核辅导员
	Status       string      `json:"status" db:"status"`               //本次审核状态
}

//show BackAppList
type ShowBackAppList struct {
	ID           string      `json:"id"`            //学号
	HealthStatus string      `json:"health_status"` //健康状态
	FromWhere    string      `json:"from_where"`    //从何地出发
	StartTime    string      `json:"start_time"`    //出发时间
	GetTime      string      `json:"get_time"`      //预计到校时间
	Vehicle      string      `json:"vehicle"`       //交通工具
	StaySchool   string      `json:"stay_school"`   //是否住校
	CarNumber    string      `json:"car_number"`    //自驾车牌
	SuiKangMini  string      `json:"suikang_mini"`  //是否填报穗康小程序
	HealthPhoto  image.Image `json:"health_photo"`  //健康码
	TripPhoto    image.Image `json:"trip_photo"`    //行程卡
	Preceptor    string      `json:"preceptor"`     //审核辅导员
	Status       string      `json:"status"`        //本次审核状态
}

type Person struct {
	ID         string `json:"id" db:"id"`                 //学号
	Name       string `json:"name" db:"name"`             //姓名
	Academy    string `json:"academy" db:"academy"`       //学院
	Profession string `json:"profession" db:"profession"` //专业
	Class      string `json:"class" db:"class"`           //班级
	Grade      string `json:"grade" db:"grade"`           //年级
}

type ShowPerson struct {
	ID         string `json:"id"`                //学号
	Page       int64  `json:"page" example:"1"`  //页号
	Size       int64  `json:"size" example:"10"` //条数
	Name       string `json:"name"`              //姓名
	Academy    string `json:"academy"`           //学院
	Profession string `json:"profession"`        //专业
	Class      string `json:"class"`             //班级
	Grade      string `json:"grade"`             //年级
}

type Auth struct {
	ID      string `json:"id" db:"id"`           //工号
	Name    string `json:"name" db:"name"`       //姓名
	Academy string `json:"academy" db:"academy"` //学院

}
