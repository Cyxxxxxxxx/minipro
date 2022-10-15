package controller

import (
	"fmt"
	"minipro/mysql"
	"minipro/service"
	"minipro/setting"
	"minipro/tool"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// JudgePermission 判断权限等级
// @Summary 判断权限等级
// @Tags 权限相关接口
// @version 0.0.1
// @Description  加载时返回权限等级(1.无权限,2.管理员,3.超级管理员)"）
// @Accept application/json
// @Produce application/json
// @Param userid query string true "用户id"
// @Success 200 {object} RespAuth
// @Router /api/judgePermission [get]
func JudgePermission(c *gin.Context) {
	//获得id
	userid := c.Query("userid")
	//检验参数
	if userid == "" {
		//controller.ResponseError(c, controller.CodeErrPermission)
		ResponseError(c, CodeErrPermission)
		return
	}
	//校验用户身份
	//校验超管身份
	authArray := strings.Split(setting.Conf.AuthID, ";")

	for _, value := range authArray {
		//如果校验到该用户，则为超管
		if value == userid {
			ResponseSuccess(c, CodeSuccess, 1)
			return
		}
	}

	//校验是否有扫码权限,若有扫码权限则返回2，若没扫码权限则返回3
	count, err := service.JudgeConfirmPermission(userid)
	if err != nil {
		zap.L().Error("判断管理应用用户权限错误", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	if count >= 1 {
		//存在权限
		ResponseSuccess(c, CodeSuccess, 2)
		return
	} else {
		//权限不足
		ResponseSuccess(c, CodeSuccess, 3)
		return
	}
}

// JudgeConfirmPermission 判断是否具有管理员权限
// @Summary 判断是否具有管理员权限
// @Tags 权限相关接口
// @version 0.0.1
// @Description  加载时判断是否具有管理员权限,若有权限进去管理员界面，无权限则显示无权限提示
// @Produce application/json
// @Param userid query string true "用户id"
// @Success 200 {object} RespSuccess
// @Router /api/JudgeConfirmPer [get]
func JudgeConfirmPermission(c *gin.Context) {
	//获得用户id
	staffid := c.Query("userid")

	//如果id为“”
	if staffid == "" {
		zap.L().Info("Err in get userid")
		ResponseError(c, CodeErrGetUser)
		return
	}

	//根据返回的个数判断是否存在权限
	count, err := service.JudgeConfirmPermission(staffid)

	if err != nil {
		zap.L().Error("判断管理应用用户权限错误", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	if count >= 1 {
		//存在权限
		ResponseSuccess(c, CodeSuccess, nil)
		return
	} else {
		//权限不足
		ResponseSuccess(c, CodeErrPermission, nil)
		return
	}
}

// QueryStaff 教职工搜索
// @Summary 教职工搜索
// @version 0.0.1
// @Tags 相关接口
// @Description  从教职工表中搜索教职工信息,状态码(1000-成功-对应搜索的数据返回)
// @Param staffID/staffName/page/count body string true "教职工工号/教职工姓名/页数/行数-后两个为int"
// @Param userid query string true "用户id"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} RespStaff
// @Router /api/QueryStaff [post]
func QueryStaff(c *gin.Context) {
	jsonMap, err := tool.GetJsonMap(c)
	if err != nil {
		zap.L().Error("获得参数错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获得body参数
	staffID := fmt.Sprint(jsonMap["staffID"])
	name := fmt.Sprint(jsonMap["staffName"])
	page := jsonMap["page"].(float64)
	count := jsonMap["count"].(float64)

	//模糊搜索对应的数据
	staffInfo, err := service.QueryStaff(staffID, name, page, count)

	if err != nil {
		zap.L().Error("查询教职工信息错误", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	num, err := service.SearchCountStaff(staffID, name)

	resp := mysql.RespStaffInfo{
		Count:     num,
		StaffInfo: staffInfo,
	}

	ResponseSuccess(c, CodeSuccess, resp)
}

// SetAuth 管理员权限设置
// @Summary 管理员权限设置
// @version 0.0.1
// @Tags 权限相关接口
// @Description  设置管理员权限
// @Param staffID/staffName/staffCampus body string true "教职工工号/教职工姓名/校区（大学城，黄埔，桂花岗）"
// @Param userid query string true "用户id"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} RespSuccess
// @Router /api/SetAuth [post]
func SetAuth(c *gin.Context) {
	jsonMap, err := tool.GetJsonMap(c)
	if err != nil {
		zap.L().Error("获得参数错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获得body里的数据
	staffId := fmt.Sprint(jsonMap["staffID"])
	staffName := fmt.Sprint(jsonMap["staffName"])
	//staffOrg := fmt.Sprint(jsonMap["staffOrg"])
	staffCampus := fmt.Sprint(jsonMap["staffCampus"])

	//service.CheckIfInsert(staffID)
	//查询是否存在该数据
	org, err := service.CheckIfInfo(staffId, staffName)
	if err != nil {
		zap.L().Error("检查信息错误", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//如果为空，则说明教职工表查无该用户,查找学生表
	var info *string
	if org.Count == 0 {
		resp, err := service.CheckIfInfoStu(staffId, staffName)
		//ResponseSuccess(c, CodeCheckNotUser, nil)
		if err != nil {
			zap.L().Error("检查本科生信息错误", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		//如果学生表也查无该数据，说明不存在该用户
		if resp.Count == 0 {
			zap.L().Info("查无该用户:" + staffId + "  " + staffName)
			ResponseSuccess(c, CodeCheckNotUser, nil)
			return
		} else {
			info = resp.O_STUDENT_BASIC_COLLEGENAME
		}
	} else {
		info = org.O_STAFF_BASIC_ORG
	}

	//处理业务
	//判断是否已经存在该用户
	count, err := service.CheckIfAuth(staffId)
	if count >= 1 {
		ResponseSuccess(c, CodeCheckUser, nil)
		return
	}
	if err != nil {
		zap.L().Error("检查权限失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	err = service.SetAuth(staffId, staffName, info, staffCampus)

	if err != nil {
		zap.L().Error("插入扫码权限表失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, CodeSuccess, nil)
}

// DeletePermission 删除管理员权限
// @Summary 删除管理员权限
// @Tags 权限相关接口
// @version 0.0.1
// @Description  删除管理员权限，从数据表中删掉该行数据
// @Param staffID body array true "辅导员工号"
// @Param userid query string true "用户id"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} RespSuccess
// @Router /api/DeleteAuth [post]
func DeleteAuth(c *gin.Context) {
	jsonMap, err := tool.GetJsonMap(c)
	if err != nil {
		zap.L().Error("获得参数错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获得工号数组
	staffArray := jsonMap["staffID"]
	value := reflect.ValueOf(staffArray)
	//如果不为数组，则参数错误
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		zap.L().Info("参数错误（需要一个数组参数）", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	for i := 0; i < value.Len(); i++ {
		err = service.DeleteAuth(value.Index(i).Interface().(string))
		if err != nil {
			zap.L().Error("删除权限表数据错误", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
	}
	ResponseSuccess(c, CodeSuccess, nil)
}

// ShowPermission 显示权限用户
// @Summary 显示权限用户
// @Tags 权限相关接口
// @version 0.0.1
// @Description  加载所有具有管理员权限的用户信息
// @Param page/count body int true "页数/行数"
// @Param userid query string true "用户id"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} RespStaff
// @Router /api/ShowStaffInfo [post]
func ShowAuth(c *gin.Context) {
	jsonMap, err := tool.GetJsonMap(c)
	if err != nil {
		zap.L().Error("获得参数错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获得body里的页数和行数
	page := jsonMap["page"].(float64)
	count := jsonMap["count"].(float64)

	//获得扫码权限用户
	staffInfo, num, err := service.ShowAuth(page, count)
	resp := mysql.RespStaffInfo{
		StaffInfo: staffInfo,
		Count:     num,
	}
	if err != nil {
		zap.L().Error("删除权限表数据错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	ResponseSuccess(c, CodeSuccess, resp)
}

// SetCampus 修改校区
// @Summary 修改校区信息,参数按照括号内填写
// @Tags 相关接口
// @version 0.0.1
// @Description  修改校区
// @Param staffID/staffCampus body string true "教职工工号/教职工校区（大学城、黄埔、桂花岗）"
// @Param userid query string true "用户id"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} RespSuccess
// @Router /api/UpdateCampus [post]
func SetCampus(c *gin.Context) {
	jsonMap, err := tool.GetJsonMap(c)
	if err != nil {
		zap.L().Error("获得参数错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	staffId := fmt.Sprint(jsonMap["staffID"])
	campus := fmt.Sprint(jsonMap["staffCampus"])

	err = service.SetCampus(staffId, campus)

	if err != nil {
		zap.L().Error("删除权限表数据错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	ResponseSuccess(c, CodeSuccess, nil)
}

// GetCollegeName 获得学院名称列表
// @Summary 获得数据库中存在的所有学院信息列表
// @Tags 返校应用相关接口
// @version 0.0.1
// @Description  获得学院名称列表
// @Param userid query string true "用户id"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} RespCollege
// @Router /api/GetCollegeName [get]
func GetCollegeName(c *gin.Context) {
	college, err := service.GetCollegeName()
	if err != nil {
		zap.L().Error("删除权限表数据错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	ResponseSuccess(c, CodeSuccess, college)
	return
}
