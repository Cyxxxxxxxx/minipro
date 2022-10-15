package routes

import (
	"minipro/controller"
	"minipro/logger"
	"minipro/middlewares"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"

	// "github.com/swaggo/gin-swagger/swaggerFiles"
	swaggerFiles "github.com/swaggo/files"
)

func Setup(mode string) *gin.Engine {
	//如果设置mode为release则设置gin为该模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middlewares.Cors())
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// r.StaticFile("/WW_verify_CSaj5WtuBq3AemhM.txt", "./WW_verify_CSaj5WtuBq3AemhM.txt")

	//管理后台
	r.GET("/Getscheduledtime/:id", controller.GetScheduledTimeHandle)
	r.POST("/create", controller.CreateNewStudentHandle)
	r.GET("/import", controller.ImportStudentMsgHandle)
	r.GET("/msg/:id", controller.GetStudentMsgById)

	//返校管理后台路由组(需要超级管理员权限)---包括权限管理、返校管理全部接口以及返校详情里的添加功能
	v := r.Group("/api").Use(middlewares.JudgePermissionWeb)
	{
		//显示所有授权扫码用户信息
		v.POST("/ShowStaffInfo", controller.ShowAuth)
		//查询用户
		v.POST("/QueryStaff", controller.QueryStaff)
		//赋予权限
		v.POST("/SetAuth", controller.SetAuth)
		//删除权限
		v.POST("/DeleteAuth", controller.DeleteAuth)
		//修改校区--权限管理界面
		v.POST("/UpdateCampus", controller.SetCampus)
		//获得学院信息
		v.GET("/GetCollegeName", controller.GetCollegeName)
	}

	//需要管理员权限
	v1 := r.Group("/api").Use(middlewares.JudgePermissionAll)
	{
		v1.POST("/InsertPersonOfStu", controller.ImportToPerson)
	}

	return r
}
