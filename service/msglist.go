package service

import (
	"fmt"
	"minipro/models"
	"minipro/mysql"
)

func GetMsgList(p *models.MsgList) ([]*models.ShowStudentMsg, error, int64) {
	if p.Type == 1 {
		return mysql.GetMsgListNotInSch(p)
	} else if p.Type == 2 {
		return mysql.GetMsgListInSch(p)
	} else if p.Type == 3 {
		return mysql.GetMsgListNotApplySch(p)
	} else if p.Type == 4 {
		return mysql.GetMsgListPastTime(p)
	} else if p.Type == 5 {
		return mysql.GetMsgListPlanTSch(p)
	}
	return nil, fmt.Errorf("查询条件有误!!!"), 0
}
