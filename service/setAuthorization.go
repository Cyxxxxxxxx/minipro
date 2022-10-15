package service

import "minipro/mysql"

type Response struct {
	Code int64       `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func JudgeConfirmPermission(userid string) (int, error) {
	//判断数据表中是否有该用户
	count, err := mysql.SelectAuth(userid)
	return count, err
}

func QueryStaff(staffid string, name string, page float64, count float64) ([]mysql.StaffInfoStruct, error) {
	if staffid == "" {
		staffid = "%"
	}
	if name == "" {
		name = "%"
	}
	staffInfo, err := mysql.SelectStaffInfo(staffid, name, page, count)
	return staffInfo, err
}

func SetAuth(staffid string, staffName string, staffOrg *string, staffCampus string) error {
	err := mysql.InsertInAuth(staffid, staffName, staffOrg, staffCampus)
	return err
}

func DeleteAuth(staffid string) error {
	err := mysql.DeleteAuth(staffid)
	return err
}

func ShowAuth(page float64, count float64) ([]mysql.StaffInfoStruct, int, error) {
	staffInfo, err := mysql.GetStaffInfo(page, count)
	if err != nil {
		return staffInfo, 0, err
	}

	number, err := mysql.SelectCountOfAuth()

	return staffInfo, number, err
}

func SetCampus(staffId string, campus string) error {
	err := mysql.UpdateCampus(staffId, campus)
	return err
}

func GetCollegeName() ([]string, error) {
	//获得数据库中所有的学院信息
	college, err := mysql.SelectCollege()
	return college, err
}

func CheckIfInfo(staffID string, staffName string) (mysql.RespOrg, error) {
	org, err := mysql.SelectIfStaff(staffID, staffName)
	if err != nil {
		return org, err
	}
	return org, err
}

func CheckIfInfoStu(staffID string, staffName string) (mysql.RespCheckCollege, error) {
	resp, err := mysql.SelectIfStu(staffID, staffName)
	return resp, err
}

func CheckIfAuth(staffID string) (int, error) {
	count, err := mysql.SelectIfAuth(staffID)
	return count, err
}

func SearchCountStaff(staffID string, staffName string) (int, error) {
	count, err := mysql.SelectCountOfSearch(staffID, staffName)
	return count, err
}
