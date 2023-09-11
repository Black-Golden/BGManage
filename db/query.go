package db

import (
	"github.com/ethereum/api-in/types"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

// 根据用户名查询密码--todo:新建管理员时密文存储
func QueryPassword(engine *xorm.Engine, payload *types.PayLoad) (string, error) {
	return "", nil
}

func GetUser(engine *xorm.Engine, payload *types.PayLoad) ([]types.Users, error) {
	var users []types.Users

	sessionSql := engine.Table("users")

	if payload.UserName != "" {
		sessionSql = sessionSql.Where("`f_userName` = ?", payload.UserName)
	}

	if payload.Uid != "" {
		sessionSql = sessionSql.Where("`f_uid` = ?", payload.Uid)
	}
	// 时间
	if payload.StartTime != "" && payload.EndTime != "" {
		sessionSql = sessionSql.Where("?<=`f_createTime`<?", payload.StartTime, payload.EndTime)
	}
	//一期不考虑分页
	err = sessionSql.Find(&users)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return users, nil
}
