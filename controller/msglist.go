package controller

import (
	"minipro/models"
	"minipro/service"
	"minipro/tool"
	"minipro/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetMsgListHandle 根据传入的标识号 获取 1- 未返校  2- 已返校  3- 未申请 的学生信息
// @Summary 展示学生返校情况
// @Description 调用用展示学生返校的几种情况	获取 1- 未返校  2- 已返校  3- 未申请 的学生信息
// @Tags 展示相关接口
// @Accept application/json
// @Produce application/json
// @Param object query models.MsgList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMsgList
// @Router /api/msglist [get]
func GetMsgListHandle(c *gin.Context) {
	// 接收参数
	p := &models.MsgList{
		Page:        1,
		Size:        10,
		Type:        0,
		StudentNum:  "%",
		StudentName: "%",
		Campus:      "%",
		CollegeName: "%",
		Grade:       "%",
		ArrivalTime: "%",
	}
	if err := c.ShouldBindQuery(p); err != nil {
		// 参数错误
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 处理字符串
	tool.SetWildCardInSta(p)
	// 业务处理

	datas, err, count := service.GetMsgList(p)
	if err != nil {
		zap.L().Error("service.GetMsgList(p) failed err : ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, CodeSuccess, gin.H{
		"data":  datas,
		"count": count,
	})

}

// GetStudentMsgById 根据学号去获取信息
// @Summary 根据学号去获取信息
// @Description 根据学号去获取信息
// @Tags 展示相关接口
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} RespSuccess
// @Router /msg/:id [get]
func GetStudentMsgById(c *gin.Context) {
	stuid := c.Param("id")
	if stuid == "" {
		ResponseError(c, CodeInvalidParam)
		return
	}
	err, data := mysql.GetStudentMsg(stuid)
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, "没有此学号！")
		return
	}

	if value, ok := data.(*models.StudentMsg); ok {
		ResponseSuccess(c, CodeSuccess, gin.H{
			"data": value,
			"tag":  "本科生",
		})
		return
	}

	value, ok := data.(*models.StudentMsgPg)
	if !ok {
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回数据
	ResponseSuccess(c, CodeSuccess, gin.H{
		"data": value,
		"tag":  "研究生或博士",
	})
}
