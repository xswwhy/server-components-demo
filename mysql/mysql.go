package mysqlOper

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	db *sql.DB
}

func NewMysql(name string, password string, addr string, dbName string) (p *Mysql, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", name, password, addr, dbName)
	p = &Mysql{}
	// 连接数据库
	p.db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("sql.Open() err:", err)
		return
	}
	// 测试连接
	err = p.db.Ping()
	if err != nil {
		fmt.Println("db.ping err:", err)
		return
	}
	p.db.SetMaxIdleConns(20) // 设置最大连接数量
	p.db.SetMaxOpenConns(20) // 设置最大闲置连接数量
	return
}

// 获取单行数据
func (m *Mysql) SelectRow() error {
	sqlstr := "select username,password,email from users where id = 13"
	row := m.db.QueryRow(sqlstr) // 读取一行数据
	var username string
	var password string
	var email string
	// 解析读取的数据，就算有多行数据也只会解析第一行数据，然后将连接放回连接池
	err := row.Scan(&username, &password, &email)
	if err != nil {
		fmt.Println("row.Scan err:", err)
		return err
	}
	fmt.Println(username, password, email)
	return nil
}

// 获取多行数据
func (m *Mysql) SelectRows() error {
	sqlstr := "select username,password,email from users where id > 13"
	rows, err := m.db.Query(sqlstr)
	if err != nil {
		fmt.Println("db.Query err:", err)
		return err
	}
	defer rows.Close() // 这个非常重要，要手动关闭连接，放回连接池

	var username string
	var password string
	var email string
	// 解析读取的多行数据,循环读取
	for rows.Next() {
		err := rows.Scan(&username, &password, &email)
		if err != nil {
			fmt.Println("rows.Scan err:", err)
			return err
		}
		fmt.Println(username, password, email)
	}
	return nil
}

// 插入数据
func (m *Mysql) InsertData() error {
	sqlstr := `insert into users(username,password,email,age) values("asd","123456","werdg@163.com",25)`
	ret, err := m.db.Exec(sqlstr)
	if err != nil {
		fmt.Println("db.Exec err:", err)
		return err
	}
	id, err := ret.LastInsertId() // 获取插入的最后一行的id
	if err != nil {
		fmt.Println("ret.LastInsertId err:", err)
		return err
	}
	fmt.Println("插入的数据id:", id)
	return nil
}

// 更新数据
func (m *Mysql) UpdateData() error {
	sqlstr := `update users set age=age-2 where id = 15`
	ret, err := m.db.Exec(sqlstr)
	if err != nil {
		fmt.Println("db.Exec err:", err)
		return err
	}
	num, err := ret.RowsAffected() // 获取更新的行数
	if err != nil {
		fmt.Println("ret.RowsAffected err:", err)
		return err
	}
	fmt.Println("更新的行数:", num)
	return nil
}

// 	// 删除数据(和更新数据写法一样)
// 	// sqlstr = "delect from orders where order_num=?"

// 数据库预处理
func (m *Mysql) PrepareOper() error {
	sqlstr := "select username,password,email from users where id > ?"
	stmt, err := m.db.Prepare(sqlstr)
	if err != nil {
		fmt.Println("db.Prepare err:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Query(0) // rows为查询结果
	//写操作也一样 stmt.Exec()
	return nil
}

// 数据库事务
func (m *Mysql) Transaction() error {
	//事务操作
	tx, err := m.db.Begin()
	if err != nil {
		fmt.Println("db.Begin err;", err)
		return err
	}

	sqlstr1 := "update users set age=age-2 where id=15"
	sqlstr2 := "update users set age=age+2 where id=17"
	_, err = m.db.Exec(sqlstr1)
	if err != nil {
		tx.Rollback() // 回滚
	}
	_, err = tx.Exec(sqlstr2)
	if err != nil {
		tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}
	fmt.Println("事务提交成功")
	return nil
}
