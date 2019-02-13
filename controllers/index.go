package controllers

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"luaalipay/library/util"
	"luaalipay/models"
)

type IndexController struct{}

var phoneM = new(models.PhoneModel)

func (this IndexController) SaveImei(c *gin.Context) {
	imei, _ := c.GetQuery("imei")
	times, _ := c.GetQuery("times")
	phoneM.SaveImei(imei, times)

	util.TipGG(c, imei)
}

func (this IndexController) SaveImei2(c *gin.Context) {
	imei, _ := c.GetQuery("imei")
	times, _ := c.GetQuery("times")
	phoneM.SaveImei2(imei, times)

	util.TipGG(c, imei)
}

func (this IndexController) GetPhone(c *gin.Context) {
	imei, _ := c.GetQuery("imei")
	times, _ := c.GetQuery("times")
	data, _ := phoneM.GetOnePhone(imei, times)

	util.TipGG(c, data.Phone.String+"|"+data.Content.String)
}

func (this IndexController) GetPhone2(c *gin.Context) {
	imei, _ := c.GetQuery("imei")
	times, _ := c.GetQuery("times")
	data, _ := phoneM.GetOnePhone2(imei, times)

	util.TipGG(c, data.Pre_phone.String)
}

func (this IndexController) GetNeedAddFriend(c *gin.Context) {
	phone, _ := c.GetQuery("phone")

	data, _ := phoneM.GetNeedAddFriend(phone)
	phoneS := ""

	for _, v := range data {
		if phoneS == "" {
			phoneS = v.Account.String
		} else {
			phoneS = phoneS + "|" + v.Account.String
		}

	}

	util.TipGG(c, phoneS)
}

func (this IndexController) UpdateNeedAddfriend(c *gin.Context) {
	isreal, _ := c.GetQuery("isreal")
	isaddfriend, _ := c.GetQuery("isaddfriend")
	phone, _ := c.GetQuery("phone")
	toid, _ := c.GetQuery("toid")

	phoneM.UpdateNeedAddfriend(isreal, isaddfriend, phone, toid)

	util.TipGG(c, "")
}

func (this IndexController) UpdateCheckRule(c *gin.Context) {
	isreal, _ := c.GetQuery("isreal")
	isaddfriend, _ := c.GetQuery("isaddfriend")
	phone, _ := c.GetQuery("phone")

	phoneM.UpdateCheckRule(isreal, isaddfriend, phone)

	util.TipGG(c, "")
}

func (this IndexController) SaveCheckPhone(c *gin.Context) {
	account, _ := c.GetQuery("account")
	isrealname, _ := c.GetQuery("isrealname")
	ischeck, _ := c.GetQuery("ischeck")

	phoneM.SaveCheckPhone(account, isrealname, ischeck)

	util.TipGG(c, "")
}
