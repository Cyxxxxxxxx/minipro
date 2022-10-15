package mysql

import (
	"fmt"
	"minipro/models"
)

//preceptor hand insert data to minipro_person
func InsertPerson(per *models.Person) error {
	sqlStr := `INSERT INTO
					minipro_person(id, name,academy,profession,class,grade)
				VALUES (?,?,?,?,?,?)`
	ret, err := db.Exec(sqlStr, per.ID, per.Name, per.Academy, per.Profession, per.Class, per.Grade)
	if err != nil {
		return err
	}
	//lastest insert ID
	_, err = ret.LastInsertId()
	if err != nil {
		return fmt.Errorf("error inserting personList: %v", err)
	}
	return nil
}

//insert student's excused list into minipro_excused
func InsertExcused(ExList *models.ExcusedList) error {
	sqlStr := `INSERT INTO
					minipro_excused(id,leave_type,
									leave_days,leave_time_start,
									leave_time_end,photo,
									leave_school,out_where01,
									out_where02,out_where03,
									campus,person_number,
									emergency_number,leave_reason,
									status,preceptor)
				VALUES(	?,?,?,?,
						?,?,?,?,
						?,?,?,?,
						?,?,?,?)`
	ret, err := db.Exec(sqlStr, ExList.ID,
		ExList.LeaveType,
		ExList.LeaveDays,
		ExList.LeaveTimeStart,
		ExList.LeabeTimeEnd,
		ExList.Photo,
		ExList.LeaveSchool,
		ExList.OutWhere01,
		ExList.OutWhere02,
		ExList.OutWhere03,
		ExList.Campus,
		ExList.PersonNumber,
		ExList.EmergencyNumber,
		ExList.LeaveReason,
		ExList.Status,
		ExList.Preceptor,
	)
	if err != nil {
		return err
	}
	//lastest insert ID
	_, err = ret.LastInsertId()
	if err != nil {
		return fmt.Errorf("error inserting excusedList: %v", err)
	}
	return nil
}

//insert healthList into minipro_health_clock
func InsertHealthClockList(HCL *models.HealthClockList) error {
	sqlStr := `	INSERT INTO minipro_health_clock(	id,gender,phone,monitor,
												preceptor,origin,stay_or_work_school,now_stay_status,
												now_stay,out_this_day,health_status,lastest_acid,
												key_personnel,family_health,hey_family,doubtful_or_diagnosed,
												touch_doubtful_or_diagnosed,half_month_touch,green_code,half_month_get
												inoculation,inoculation_number,inoculation_date,inoculation_maker
												pramise,clock_date)
				VALUES(	?,?,?,?,
						?,?,?,?,
						?,?,?,?,
						?,?,?,?,
						?,?,?,?,
						?,?,?,?,
						?,?)`
	ret, err := db.Exec(sqlStr, HCL.ID, HCL.Gender, HCL.Phone, HCL.Monitor,
		HCL.Preceptor, HCL.Origin, HCL.StayOrWorkSchool, HCL.NowStayStatus,
		HCL.NowStay, HCL.OutThisDay, HCL.HealthStatus, HCL.LastestAcid,
		HCL.KeyPersonnel, HCL.FamilyHealth, HCL.KeyFamily, HCL.DoubtfulOrDiagnosed,
		HCL.TouchDoubtfulOrDiagnosed, HCL.HalfMonthTouch, HCL.GreenCode, HCL.HalfMonthGet,
		HCL.Inoculation, HCL.InoculationNumber, HCL.InoculationDate, HCL.InoculationMaker,
		HCL.Pramise, HCL.ClockDate)
	if err != nil {
		return err
	}
	//lastest insert ID
	_, err = ret.LastInsertId()
	if err != nil {
		return fmt.Errorf("error inserting	HealthList: %v", err)
	}
	return nil
}

//insert student's backtoschoolapplication into minipro_back_app
func InsertBackAppList(BAL *models.BackAppList) error {
	sqlStr := `	INSERT INTO minipro_back_app(
											id,health_status,from_where,start_time,
											get_time,vehicle,stay_school,car_number,
											suikang_mini,health_photo,trip_photo,preceptor,
											status)
				VALUES(
						?,?,?,?,
						?,?,?,?,
						?,?,?,?,
						?)`
	ret, err := db.Exec(sqlStr, BAL.ID, BAL.HealthStatus, BAL.FromWhere, BAL.StartTime,
		BAL.GetTime, BAL.Vehicle, BAL.StaySchool, BAL.CarNumber,
		BAL.SuiKangMini, BAL.HealthPhoto, BAL.TripPhoto, BAL.Preceptor,
		BAL.Status)
	if err != nil {
		return err
	}
	//lastest insert ID
	_, err = ret.LastInsertId()
	if err != nil {
		return fmt.Errorf("error inserting BackAppList: %v", err)
	}
	return nil
}

//query HealthClockHistory from minipro_health_clock
func GetAllDataAboutOne(p *models.ShowHealthClockList) (RetData []*models.ShowHealthClockList, err error, count int) {
	sqlFindAll := `SELECT * FROM 
						(
							SELECT  p.id as hc.id,
									p.name,
									P.academy,
									p.profession,
									p.class,
									p.grade,
									hc.gender,
									hc.phone,
									hc.monitor,
									hc.preceptor,
									hc.origin,
									hc.stay_or_work_school,
									hc.now_stay_status,
									hc.now_stay,
									hc.out_this_day,
									hc.health_status,
									hc.lastest_acid,
									hc.key_personnel,
									hc.family_health,
									hc.key_family,
									hc.doubtful_or_diagnosed,
									hc.touch_doubtful_or_diagnosed,
									hc.half_month_touch,
									hc.green_code,
									hc.half_month_get,
									hc.inoculation,
									hc.inoculation_number,
									hc.inoculation_date,
									hc.inoculation_maker,
									hc.pramise,
									hc.clock_date
							FROM minipro_person p
								right JOIN minipro_health_clock hc
									on p.id = hc.id
						)
							as sec
					WHERE 	hc.id like ?
						and hc.name like ?
						and hc.preceptor like ?
						and DATE_FORMAT(hc.clock_date,'%Y-%m-%d') like ?
					limit ? offset ?;
					`
	sqlCount := `SELECT count(*) from minipro_person p
						right JOIN minipro_health_clock hc
							on p.id = hc.id
					WHERE 	hc.id like ?
						and hc.name like ?
						and hc.preceptor like ?
						and DATE_FORMAT(hc.clock_date,'%Y-%m-%d') like ?;
	`
	start := (p.Page - 1) * p.Size
	err = db.Select(&RetData, sqlFindAll, p.ID, p.Name, p.Preceptor, p.ClockDate, p.Size, start)
	err = db.Get(&count, sqlCount, p.ID, p.Name, p.Preceptor, p.ClockDate)
	return
}

//query NOtHealthCLock student list from minipro_health_clock and minipro_person
func GetNotHealthCLock(p *models.ShowHealthClockList) (RetData []*models.Person, err error, count int) {
	sqlFindNotAll := `SELECT    p.id,
								p.name,
								P.academy,
								p.profession,
								p.class,
								p.grade
							FROM minipro_person p;
						`
	sqlNotClock := `SELECT  	hc.id	
						from minipro_health_clock hc
					WHERE 	hc.id like ?
						and hc.name like ?
						and hc.preceptor like ?
						and DATE_FORMAT(hc.clock_date,'%Y-%m-%d') like ?;
	`
	var GetNotClock []interface{}
	err = db.Select(&RetData, sqlFindNotAll)
	err = db.Select(&GetNotClock, sqlNotClock, p.ID, p.Name, p.Preceptor, p.ClockDate)

	TampData := RetData
	for k, v := range TampData {
		for _, b := range GetNotClock {
			if v.ID == b {
				RetData = append(RetData[:k-1], RetData[k+1:]...)
			}
		}
	}
	count = len(RetData)
	return
}

//quuery NotStaySchool list of student from minipro_person and minipro_back_app
func GetNotStaySchool() {

}
