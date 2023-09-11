package api

import (
	"github.com/ethereum/api-in/db"
	"github.com/ethereum/api-in/types"
	"github.com/ethereum/api-in/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"net/http"
)

//func (a *ApiService) init(c *gin.Context) {
//	buf := make([]byte, 2048)
//	n, _ := c.Request.Body.Read(buf)
//	data1 := string(buf[0:n])
//	res := types.HttpRes{}
//
//	isValid := gjson.Valid(data1)
//	if isValid == false {
//		logrus.Error("Not valid json")
//		res.Code = http.StatusBadRequest
//		res.Message = "Not valid json"
//		c.SecureJSON(http.StatusBadRequest, res)
//		return
//	}
//	name := gjson.Get(data1, "name")
//	apiKey := gjson.Get(data1, "apiKey")
//	apiSecret := gjson.Get(data1, "apiSecret")
//
//	mechanismData := types.Mechanism{
//		Name:      name.String(),
//		ApiKey:    apiKey.String(),
//		ApiSecret: apiSecret.String(),
//	}
//
//	err := a.db.CommitWithSession(a.db, func(s *xorm.Session) error {
//		if err := a.db.InsertMechanism(s, &mechanismData); err != nil {
//			logrus.Errorf("insert  InsertMechanism task error:%v tasks:[%v]", err, mechanismData)
//			return err
//		}
//		return nil
//	})
//	if err != nil {
//		logrus.Error(err)
//	}
//
//	res.Code = 0
//	res.Message = err.Error()
//	res.Data = ""
//
//	c.SecureJSON(http.StatusOK, res)
//}

func (a *ApiService) order(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])
	res := types.HttpRes{}

	isValid := gjson.Valid(data1)
	if isValid == false {
		logrus.Error("Not valid json")
		res.Code = http.StatusBadRequest
		res.Message = "Not valid json"
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	//下面将信息存入db
	res.Code = 0
	res.Message = "success"
	res.Data = "null"

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) enroll(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])
	res := types.HttpRes{}

	isValid := gjson.Valid(data1)
	if isValid == false {
		logrus.Error("Not valid json")
		res.Code = http.StatusBadRequest
		res.Message = "Not valid json"
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	uid := gjson.Get(data1, "uid")
	password := gjson.Get(data1, "password")

	user := types.Users{
		Uid:      uid.String(),
		Password: password.String(),
	}

	db.InsertUser(a.dbEngine, &user)
	//下面将信息存入db
	res.Code = 0
	res.Message = "success"
	res.Data = "null"

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) login(c *gin.Context) {
	var payload *types.LoginInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		logrus.Error(err)
		res := util.ResponseMsg(-1, "fail", err.Error())
		c.SecureJSON(http.StatusOK, res)
		return
	}
	var passWord string
	// 获取数据库中的密码
	err, has := db.QueryPassword(a.dbEngine, payload.UserName)
	if err != nil {
		res := util.ResponseMsg(-1, "fail", "User does not exist.")
		c.SecureJSON(http.StatusOK, res)
		return
	}
	if has != nil {
		passWord = has.Password
	}
	//todo: 密文比较
	if payload.Password != passWord {
		res := util.ResponseMsg(-1, "fail", "Incorrect password.")
		c.SecureJSON(http.StatusOK, res)
		return
	}

	res := util.ResponseMsg(0, "success", body)
	c.SecureJSON(http.StatusOK, res)
	return
}

func (a *ApiService) newAdmin(c *gin.Context) {
	var admin *types.Admin

	err := c.BindJSON(&admin)
	if err != nil {
		res := util.ResponseMsg(-1, "fail", err)
		c.SecureJSON(http.StatusOK, res)
		return
	}

	user := types.Admin{
		UserName: admin.UserName,
		Role:     admin.Role,
		Password: admin.Password,
	}

	err = db.InsertAdmin(a.dbEngine, &user)
	if err != nil {
		res := util.ResponseMsg(-1, "fail", err)
		c.SecureJSON(http.StatusOK, res)
		return
	}
	res := util.ResponseMsg(0, "success", nil)
	c.SecureJSON(http.StatusOK, res)
	return
}

func (a *ApiService) editAdmin(c *gin.Context) {
	var admin *types.Admin

	err := c.BindJSON(&admin)
	if err != nil {
		res := util.ResponseMsg(-1, "fail", err)
		c.SecureJSON(http.StatusOK, res)
		return
	}

	user := types.Admin{
		UserName: admin.UserName,
		Role:     admin.Role,
		Password: admin.Password,
	}

	err = db.UpdateAdmin(a.dbEngine, &user)
	if err != nil {
		res := util.ResponseMsg(-1, "fail", err)
		c.SecureJSON(http.StatusOK, res)
		return
	}
	res := util.ResponseMsg(0, "success", nil)
	c.SecureJSON(http.StatusOK, res)
	return
}

func (a *ApiService) newRole(c *gin.Context) {
	var role *types.Role

	err := c.BindJSON(&role)
	if err != nil {
		res := util.ResponseMsg(-1, "fail", err)
		c.SecureJSON(http.StatusOK, res)
		return
	}

	err = db.InsertRole(a.dbEngine, role)
	if err != nil {
		res := util.ResponseMsg(-1, "fail", err)
		c.SecureJSON(http.StatusOK, res)
		return
	}
	res := util.ResponseMsg(0, "success", nil)
	c.SecureJSON(http.StatusOK, res)
	return
}

func (a *ApiService) editRole(c *gin.Context) {
	var role *types.Role

	err := c.BindJSON(&role)
	if err != nil {
		res := util.ResponseMsg(-1, "fail", err)
		c.SecureJSON(http.StatusOK, res)
		return
	}

	err = db.UpdateRole(a.dbEngine, role)
	if err != nil {
		res := util.ResponseMsg(-1, "fail", err)
		c.SecureJSON(http.StatusOK, res)
		return
	}
	res := util.ResponseMsg(0, "success", nil)
	c.SecureJSON(http.StatusOK, res)
	return
}
