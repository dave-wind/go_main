package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	// _ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

// go mod init name
// go get xxx
// go run xxx

// mysql 数据库驱动

/*
准备连接数据库

	要想连接到SQL数据库，首先需要加载目标数据库的驱动，
	驱动里面包含着与该数据库交互的逻辑。
	sql.Open()
	数据库驱动的名称数据源名称
	o得到一个指向sqlDB这个struct的指针
	sql.DB是用来操作数据库的，它代表了0个或者多个底层
	连接的池，这些连接由sql包来维护，sql包会自动的创建和释放这些连接
*
*/

/* 小记
*   Open()函数并不会连接数据库，甚至不会验证其参数。它只是
	把后续连接到数据库所必需的 structs 给设置好了
	，而真正的连接是在被需要的时候才进行懒设置的

	sql.DB 不需要进行关闭(当然你想关闭也是可以的)
	。 它就是用来处理数据库的，而不是实际的连接
	这个抽象包含了数据库连接的池，而且会对此进行维护
	· 在使用 sql.DB 的时候，可以定义它的全局变量进行使用，也可
	以将它传递函数/方法里。
*/

/*
如何 获取驱动

正常的做法是使用sqlRegister()函数数据库驱动的名称和一个实现了 driverDriver接口的struct，来注册数据库的驱动
例如:
sql.Register("sqlserver",&drv{})

但是我们之前的例子却没写这句话，为什么?
。因为SqlServer的驱动，是在这个包被引入的时候进行了自我注册:
_ "github.com/go-sql-driver/mysql"

*/

/*

驱动自我注册

	·当go-mssqldb包被引入到时候，它的init函数将会运行
	并进行自我注册(在Go语言里，每个包的init函数都会在自动的调用)
	·在引入go-mssqldb包的时候，把该包的名设置为下划线
	，这是因为我们不直接使用数据库驱动(只需要它起的“副作用”)，我们只使用 database/sql
	o这样，如果未来升级驱动，也无需改变代码
	·Go语言没有提供官方的数据库驱动，所有的数据库驱动都
	是第三方驱动，但是它们都遵循sqldriver包里面定义的接口
*/

/*
安装数据库驱动

	这是安装Microsoft SQL Server数据库驱动的例子:
	go get github.com/denisenkom/go-mssqldb
*/

// sql包DB 是用来操作数据库的
var DB *sql.DB

const (
	IP       = "127.0.0.1"
	PORT     = "3306"
	USERNAME = "root"
	PASSWORD = "1234qwer"
	dbName   = "dave"
)

type User struct {
	id       int64
	name     string
	password string
	age      int8
	sex      int8
	phone    string
}

func main() {
	InitDB()

	// user := User{
	// 	name:     "dave6",
	// 	password: "qq",
	// }

	// 插入
	// InsertUser(user)

	// user = User{
	// 	id:       3,
	// 	name:     "hxr",
	// 	password: "1996",
	// }

	// // 删除
	// DeleteUserById(4)

	// 更新
	// UpdateUser(user)

	// 查询
	getUser := SelectUserById(3)
	fmt.Println("name", getUser.name)
	fmt.Println("password", getUser.password)
	fmt.Println("id", getUser.id)

	// 多个查询
	for _, user := range SelectAllUserId() {
		fmt.Println("selectAll----" + user.name + " " + user.password)
	}
}

func InitDB() {
	// // 数据库连接字符串
	// connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	// 	server, user, password, port, database)

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", IP, ":", PORT, ")/", dbName, "?charset=utf8"}, "")

	// root:1234qwer@tcp(127.0.0.1:3306)/dave?charset=utf8

	// fmt.Println("path==", path)

	// 这里不可以重新 := 定义变量
	DB, _ = sql.Open("mysql", path)

	if err := DB.Ping(); err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println("Connected-------------------------")

	// // 创建顶层作用域
	// ctx := context.Background()

	// // ping
	// err := DB.PingContext(ctx)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
}

func InsertUser(user User) bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//	准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`name`, `password`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	// 将参数传递到sql语句中执行
	res, err := stmt.Exec(user.name, user.password)

	if err != nil {
		fmt.Println("Exec fail")
		return false
	}

	// 	将事务提交
	tx.Commit()
	fmt.Println(res.LastInsertId())

	return true
}

// 删除

func DeleteUser(user User) bool {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("DeleteUser fail")
	}
	// id = ? 占位符
	stmt, err := tx.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		fmt.Println("DELETE Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.id)
	if err != nil {
		fmt.Println("Delete Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}

// 更新
func UpdateUser(user User) bool {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("UpdateUser tx fail")
	}
	stmt, err := tx.Prepare("UPDATE user SET name = ?, password = ? WHERE id = ?")

	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	res, err := stmt.Exec(user.name, user.password, user.id)

	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}

// 单个查询
func SelectUserById(id int) User {
	var user User
	err := DB.QueryRow("SELECT id,name,password FROM user WHERE id = ?", id).Scan(&user.id, &user.name, &user.password)
	if err != nil {
		fmt.Println("查询出错了", err.Error())
	}
	return user
}

// 全查询
func SelectAllUserId() []User {
	rows, err := DB.Query("SELECT id,name,password from user")
	if err != nil {
		fmt.Println("query all fail")
	}
	var users []User

	// 循环读取结果
	for rows.Next() {
		// 定义一个变量 传进去, 改变指针
		var user User
		err := rows.Scan(&user.id, &user.name, &user.password)
		if err != nil {
			fmt.Println("for row fail")
		}
		users = append(users, user)
	}
	return users
}

func DeleteUserById(id int) bool {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("DeleteUser fail")
	}
	// id = ? 占位符
	stmt, err := tx.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		fmt.Println("DELETE Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("Delete Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}
