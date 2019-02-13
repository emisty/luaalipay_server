package models

import (
	"database/sql"
	// "fmt"
	"luaalipay/library/mysql"
	// "strconv"
	// "time"
)

type Phone struct {
	Imei      sql.NullString
	Phone     sql.NullString
	Sendid    sql.NullString
	Content   sql.NullString
	Account   sql.NullString
	Pre_phone sql.NullString
}

type PhoneModel struct{}

func (m PhoneModel) SaveImei(imei, times string) (err error) {
	getDb := mysql.GetDB()
	insert_sql2 := `INSERT INTO phone (imei,createtime) VALUES ('` + imei + `','` + times + `')`
	_, err = getDb.Exec(insert_sql2)

	return err
}

func (m PhoneModel) SaveImei2(imei, times string) (err error) {
	getDb := mysql.GetDB()
	insert_sql2 := `INSERT INTO checkrule (imei,createtime) VALUES ('` + imei + `','` + times + `')`
	_, err = getDb.Exec(insert_sql2)

	return err
}

func (m PhoneModel) GetOnePhone(imei, times string) (data Phone, err error) {
	sql := `select 
	  cv.phone,
	  cv.sendid,
	  af.content
	from
	  phone cv 
	  left join addfriendmes af 
	    on af.mesid = cv.sendid
	where cv.imei = "` + imei + `" and cv.createtime = '` + times + `'`
	err = mysql.GetDB().SelectOne(&data, sql)
	return data, err
}

func (m PhoneModel) GetOnePhone2(imei, times string) (data Phone, err error) {
	sql := `select 
	  cv.pre_phone
	from
	  checkrule cv 
	where cv.imei = "` + imei + `" and cv.createtime = '` + times + `'`
	err = mysql.GetDB().SelectOne(&data, sql)
	return data, err
}

func (m PhoneModel) GetNeedAddFriend(phone string) (data []Phone, err error) {
	_, err = mysql.GetDB().Select(&data, `SELECT 
	  account 
	FROM
	  needaddfriend 
	WHERE toid = "`+phone+`" 
	  AND ischeck = 0 
	LIMIT 0, 10 `)

	return data, err
}

func (m PhoneModel) UpdateNeedAddfriend(isreal, isaddfriend, phone, toid string) (err error) {
	getDb := mysql.GetDB()

	sql := `UPDATE needaddfriend SET isaddfriend=` + isaddfriend + `,isrealname=` + isreal + `,ischeck= 1 WHERE toid='` + toid + `' and account = '` + phone + `'`

	_, err = getDb.Exec(sql)

	return err
}

func (m PhoneModel) UpdateCheckRule(nowadd, pre, times string) (err error) {
	getDb := mysql.GetDB()

	sql := `UPDATE checkrule SET nowadd=` + nowadd + ` WHERE pre_phone = '` + pre + `' and createtime = '` + times + `'`

	_, err = getDb.Exec(sql)

	return err
}

func (m PhoneModel) SaveCheckPhone(account, isrealname, ischeck string) (err error) {
	getDb := mysql.GetDB()
	insert_sql2 := `INSERT INTO needcheckaccount (account,isrealname,ischeck) VALUES ('` + account + `','` + isrealname + `','` + ischeck + `')`
	_, err = getDb.Exec(insert_sql2)

	return err
}
