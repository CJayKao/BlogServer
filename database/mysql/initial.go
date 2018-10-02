package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

var DB *sql.DB

func MysqlInit() {
	DB, err := sqlx.Connect("mysql", "123456")
	if err != nil {
		//logger.Fatal(err, "mysql Connect Fail")
		//log.Fatalln(err)
	}
	DB.SetMaxOpenConns(0)
	DB.SetMaxIdleConns(1000)
	DB.SetConnMaxLifetime(0)
}

type MysqlStorager struct {
	User    User
	Fllower Fllower
}

type User struct {
}

type Fllower struct {
}

func functionConn() *MysqlStorager {
	return &MysqlStorager{
		User:    User{},
		Fllower: Fllower{},
	}
}
