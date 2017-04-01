package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Test() {
	users := make([]User, 0)
	db, err := sql.Open("mysql", "tongna:tongna@tcp(192.168.0.199:3306)/24hours")
	if err == nil {
		rows, err := db.Query("select * from user")
		if err != nil {
			fmt.Print(err)
		}
		// var users []User
		for rows.Next() {
			rows.Columns()

			user := User{}
			err = rows.Scan(&user.Id, &user.Name, &user.Job, &user.Age, &user.Sex)
			if err != nil {
				fmt.Print(err)
			}
			users=append(users,user)
		}
	}
         for v,x:=range users{
		 fmt.Println(v)
		 fmt.Println(x)
	 }
}
func xx(db *sql.DB, err error) {
	rows, queryerr := db.Query("SELECT id,name,username,logintimes FROM user_info ")
	checkErr(queryerr)
	for rows.Next() {
		var idx int64
		var name string
		var username string
		var logintimes int
		rows.Columns()
		err = rows.Scan(&idx, &name, &username, &logintimes)
		checkErr(err)
		fmt.Println(idx, name, username, logintimes)
	}
}
func checkErr(erro error) {
	//fmt.Println(erro.Error())
}
type User struct {
	Id int
	Name   string
	Job    string
	Age    int
	Sex    bool
}