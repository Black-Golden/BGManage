package db

import (
	"fmt"
	"github.com/ethereum/api-in/types"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"log"
)

// 插入
func InsertUser(engine *xorm.Engine, user *types.Users) {
	rows, err := engine.Table("users").Insert(user)
	if err != nil {
		log.Println(err)
		return
	}
	if rows == 0 {
		fmt.Println("插入失败")
		return
	}
	fmt.Println("插入成功")
}

func InsertAdmin(engine *xorm.Engine, admin *types.Admin) error {
	rows, err := engine.Table("admins").Insert(admin)
	if err != nil {
		log.Println(err)
		return err
	}
	if rows == 0 {
		fmt.Println("插入失败")
		return errors.New("insert null")
	}
	fmt.Println("插入成功")
	return nil
}

func InsertRole(engine *xorm.Engine, role *types.Role) error {
	rows, err := engine.Table("role").Insert(role)
	if err != nil {
		log.Println(err)
		return err
	}
	if rows == 0 {
		fmt.Println("插入失败")
		return errors.New("insert null")
	}
	fmt.Println("插入成功")
	return nil
}

func UpdateAdmin(engine *xorm.Engine, admin *types.Admin) error {
	_, err := engine.Table("admins").Where("f_userName=?", admin.UserName).Update(admin)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func UpdateRole(engine *xorm.Engine, role *types.Role) error {
	_, err := engine.Table("role").Where("f_roleName=?", role.RoleName).Update(role)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

/*
// 删除
func deleteUser(engine *xorm.Engine, name string) {
	user := types.Users{Username: name}
	rows, err := engine.Delete(&user)
	if err != nil {
		log.Println(err)
		return
	}
	if rows == 0 {
		fmt.Println("删除失败")
		return
	}
	fmt.Println("删除成功")
}

// 修改
func UpdateUser(engine *xorm.Engine, user *types.Users) {
	//Update(bean interface{}, condiBeans ...interface{}) bean是需要更新的bean,condiBeans是条件
	update, err := engine.Update(user, types.Users{Id: user.Id})
	if err != nil {
		log.Println(err)
		return
	}
	if update > 0 {
		fmt.Println("更新成功")
		return
	}
	log.Println("更新失败")
}


// 事务
func sessionUserTest(engine *xorm.Engine, user *types.Users) {
	session := engine.NewSession()
	session.Begin()
	_, err := session.Insert(user)
	if err != nil {
		session.Rollback()
		log.Fatal(err)
	}
	user.Username = "mac"
	_, err = session.Update(user, types.Users{Id: user.Id})
	if err != nil {
		session.Rollback()
		log.Fatal(err)
	}
	err = session.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
*/
