package controller

import (
	"errors"
	"fmt"
	"minipro/models"
	"minipro/mysql"
	"minipro/service"
	"minipro/tool"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type InvokeMsg struct {
	Id     string `form:"id"`     // 学生学号
	Userid string `form:"userid"` // 扫码人id
}

// DataShowHandle 扫码后接收到学生学号显示学生信息
// @Summary 显示学生信息
// @Description 扫码后接收到学生学号显示学生信息接口
// @Tags 扫码相关接口
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseDataShow
// @Router /api/datashow/:id [get]
func DataShowHandle(c *gin.Context) {
	// 接收参数
	stuid := c.Param("id")

	postid, err := strconv.ParseInt(stuid, 10, 64)
	if err != nil {
		zap.L().Error("可能未刷新 传入的学号为错误值 ： ", zap.Any("stuid = ", stuid))
		ResponseError(c, CodeScanAgain)
		return
	}
	stuid = strconv.FormatInt(postid, 10)

	// 业务处理(传入学生id）
	err, data := service.DataShow(stuid)
	if errors.Is(err, mysql.ErrorAlreadyInSch) {
		// 已经完成返校
		ResponseSuccess(c, CodeAlreadyInSch, nil)
		return
	} else if errors.Is(err, mysql.ErrorNotApply) {
		// 未申请返校
		ResponseSuccess(c, CodeNotApply, nil)
		return
	} else if errors.Is(err, mysql.ErrorPassDate) {
		// 申请已逾期
		ResponseSuccess(c, CodePassDate, nil)
		return
	} else if errors.Is(err, mysql.ErrorFrontData) {
		ResponseSuccess(c, CodeErrorFrontData, nil)
		return
	}

	if err != nil {
		// 说明出现了逾期之外的错误返回服务繁忙
		ResponseError(c, CodeServerBusy)
		return
	}

	if value, ok := data.(*models.StudentMsg); ok {
		ResponseSuccess(c, CodeSuccess, value)
		return
	}

	value, ok := data.(*models.StudentMsgPg)
	if !ok {
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回数据
	ResponseSuccess(c, CodeSuccess, value)
}

// DataUpdateHandle 扫码后示确认学生返校的确认接口
// @Summary 修改学生返校状态
// @Description 扫码后示确认学生返校的确认接口
// @Tags 扫码相关接口
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseUpdateDate
// @Router /api/dataupdate [get]
func DataUpdateHandle(c *gin.Context) {
	// 接收参数
	p := &InvokeMsg{}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("ShouldBindQuery(p) err = ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
	}

	// 业务处理(传入学生id 与 扫码人id）
	err := service.DataUpdate(p.Id, p.Userid) // 欠缺扫码人id
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ok, err := tool.UpdateBTSstatus(p.Id, "22", false)
	if !ok || err != nil {
		err = service.UnDoDataUpdate(p.Id, p.Userid)
		if err != nil {
			zap.L().Error("!!!!!!重大错误！！！！数据不一致，撤销操作失败！ err  = ", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		zap.L().Error(" 更新操作失败，撤销操作成功  err = ", zap.Error(err))
		ResponseError(c, CodeCanNotUpdateXYM)
		return
	}
	ResponseSuccess(c, CodeBackSuccess, nil)
}

// GetScheduledTimeHandle 获取学生返校状态对应的描述
func GetScheduledTimeHandle(c *gin.Context) {
	// 接收参数
	stuid := c.Param("id")
	data, err := service.GetScheduledTime(stuid)
	if err != nil {
		zap.L().Error("service.GetScheduledTime(stuid) failed err = ", zap.Error(err))
		ResponseErrorWithMsg(c, CodeNoThisStudentOrNotApply, "错误的学号或该学生暂未申请返校")
		return
	}
	// 刷新状态
	msg := ""
	// 可用map 写更好看而不需要一个个判断（ 或者一个函数里面switch 去映射）
	if data.ArrivalStatus == models.StatusNotInSch {
		msg = fmt.Sprintf("返校计划 ： %s", data.ExpectedBackTime.Format("2006-01-02"))
	} else if data.ArrivalStatus == models.StatusInSch {
		msg = "已返校"
	} else if data.ArrivalStatus == models.StatusLateToSch {
		msg = "申请已逾期,请重新申请"
	} else if data.ArrivalStatus == models.StatusDelayToSch {
		msg = "延缓返校"
	}

	ResponseSuccess(c, CodeSuccess, msg)

}

// CreateNewStudentHandle 新增未返校学生
// @Summary 新增未返校学生
// @Tags 新增未返校学生
// @version 0.0.1
// @Description  新增未返校的学生
// @Accept application/json
// @Produce application/json
// @Param object query models.CreateStudent true "参数信息"
// @Success 200 {object} _ResponseUpdateDate
// @Router /create [post]
func CreateNewStudentHandle(c *gin.Context) {

	p := &models.CreateStudent{}
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error("  c.ShouldBindJSON(p) failed err = ", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, service.removeTopStruct(errs.Translate(trans)))
		return
	}

	//syncData.SyncStuFx2App()

	err := service.CreateNewStudent(p)
	if err != nil {
		zap.L().Error(" service.CreateNewStudent(p) failed err = ", zap.Error(err))
		if errors.Is(err, models.DontInsertAgain) {
			ResponseError(c, CodeDontInsertAgain)
			return
		}
		ResponseError(c, CodeInsertErr)
		return
	}

	ok, err := tool.UpdateBTSstatus(p.StudentId, "20", true)
	if !ok || err != nil {
		zap.L().Error(" tool.UpdateBTSstatus(p.StudentId,20) failed err = ", zap.Error(err))
		err = service.UnDoDataCreate(p)
		if err != nil {
			zap.L().Error("!!!!!!重大错误！！！！数据不一致，撤销操作失败！ err  = ", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		ResponseError(c, CodeUpdataXYMFail)
		return
	}
	ResponseSuccess(c, CodeSuccess, nil)
}
