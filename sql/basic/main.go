package main

import (
	"database/sql"
	"fmt"

	"github.com/GerryLon/go-toolkit/utils"
	_ "github.com/go-sql-driver/mysql"
)

// go操作mysql基础代码
/* sql:
drop database if exists `dbtest`;
create database dbtest default charset utf8 collate utf8_general_ci;

use `dbtest`;

create table if not exists `user` (
	`id` int not null auto_increment,
	`name` varchar(20) not null,
	`age` int default 0,
	primary key(`id`)
) engine=InnoDB default charset=utf8 auto_increment=1;
*/
func main() {
	db := openDB("mysql", getDSN())
	defer db.Close()

	// (&sync.Once{}).Do(func() {
	// 插入100条测试数据
	stmt, err := db.Prepare(`INSERT INTO user(name, age) VALUES(?, ?)`)
	handleErr(err)
	for i := 0; i < 100; i++ {
		_, err := stmt.Exec(utils.RandStr(8, 8), i+1)
		handleErr(err)
	}
	// })

	// 寻找名字中带"="号的
	rows, err := db.Query(`SELECT id, name, age FROM user WHERE name LIKE '%=%' ORDER BY AGE ASC`)
	handleErr(err)
	users := table2object(rows)
	fmt.Printf("%v\n", users)
}

// 实际中可以通过配置文件读取
func getDSN() string {
	dsn := "root:gerrylon@tcp(127.0.0.1:3306)/dbtest?charset=utf8"
	return dsn
}

func openDB(driverName, dsn string) *sql.DB {
	db, err := sql.Open(driverName, dsn)
	handleErr(err)
	return db
}

// 用户对象
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("[User: id=%d, name=%s, age=%d]\n", u.Id, u.Name, u.Age)
}

// 数据表字段转换为实体
// 如果要想做得完美, 可以通过反射来完成
// elem := reflect.ValueOf(&u).Elem()
// field := elem.FieldByName("fieldName")
// field.Set("value")
// rows.Columns()
// rows.ColumnTypes()
// rows.Scan()
func table2object(rows *sql.Rows) []User {
	objects := make([]User, 0)
	for rows.Next() {
		u := User{}
		rows.Scan(&u.Id, &u.Name, &u.Age)
		objects = append(objects, u)
	}
	return objects
}

// 示例代码, 实际中需要根据情况处理
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
