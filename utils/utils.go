package utils

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql() {

	fmt.Println(".... init mysql")

	driverName := beego.AppConfig.String("driveName")

	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1

		CreateTableWithUser()

		CreateTableWithArtcle()

	}

}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
			id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
			username VARCHAR(64),
			password VARCHAR(64),
			status INT(4),
			createtime INT(10)
			);`
	ModifyDB(sql)
}

//创建文章表

func CreateTableWithArtcle() {
	sql := `CREATE TABLE IF NOT EXISTS article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
