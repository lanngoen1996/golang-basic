package mock

import (
	"database/sql"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lanngoen1996/golang-basic/core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MockConn() (sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Printf("an error '%s' was not expected when opening a stub database connection\n", err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	conn, _ := gorm.Open(dialector, &gorm.Config{})

	core.DB = conn

	return mock, db
}
