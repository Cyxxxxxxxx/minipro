package controller

import (
	"fmt"
	"minipro/tool"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type FailMessage struct {
	StudentId string `json:"student_id"`
	FailMsg   error  `json:"fail_msg"`
}

// ImportStudentMsgHandle 读取excel文档导入学生信息表
// @Summary 读取excel文档导入学生信息表
// @Tags 信息导入导出相关接口
// @version 0.0.1
// @Description  读取excel文档导入学生信息表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} _ResponseData
// @Router /ImportToPerson [get]
func ImportToPerson(c *gin.Context) {
	res, err := tool.ReadExcel(viper.GetString("filename"))
	if err != nil {
		ResponseError(c, CodeFileErr)
		return
	}

	cntsucc, cntfail := 0, 0
	failList := make([]FailMessage, 0)
	for _, row := range res {
		for key, value := range row {
			if key == viper.GetString("indexname") {
				ok, err := tool.UpdateBTSstatus(value, "20", true)
				if !ok || err != nil {
					fail := FailMessage{
						StudentId: value,
						FailMsg:   err,
					}
					zap.L().Error("该学号导入失败 -------", zap.String("学号为 ：", value))
					zap.L().Error("错误原因 ： ", zap.Error(err))
					failList = append(failList, fail)
					cntfail++
				} else {
					cntsucc++
				}
			}
		}
	}

	if len(failList) <= 0 {
		ResponseSuccess(c, CodeSuccess, fmt.Sprintf("导入学生信息表成功%d条,全部导入成功！", cntsucc))
		return
	}

	ResponseSuccess(c, CodeFilePart, gin.H{
		"message": fmt.Sprintf("导入学生信息表成功%d条，失败%d条", cntsucc, cntfail),
		"list":    failList,
	})

}
