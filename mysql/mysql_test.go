package mysqlOper

import (
	"fmt"
	"log"
	"testing"
)

func TestOperatorMysql(t *testing.T) {
	mysql, err := NewMysql("root", "12345678", "192.168.0.126", "test")
	if err != nil {
		log.Fatal("mysql创建失败")
	}

	err = mysql.SelectRow()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("单行数据测试结束-----------------------")

	err = mysql.SelectRows()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("多行数据测试结束-----------------------")

	err = mysql.InsertData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("插入数据测试结束-----------------------")

	err = mysql.UpdateData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("更新数据测试结束-----------------------")

	err = mysql.Transaction()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("提交事务测试结束-----------------------")

}
