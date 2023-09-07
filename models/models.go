package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sony/sonyflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Sf *sonyflake.Sonyflake
	Db *gorm.DB
)

type PaginateContent struct {
	Count                  int64
	CurrentPage, TotalPage int
	Contents               interface{}
	Pages                  []int
}

func init() {
	var st sonyflake.Settings
	var err error

	Sf = sonyflake.NewSonyflake(st)

	if Sf == nil {
		panic("sonyflake not created")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			// IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			// ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful: true, // Disable color
		},
	)

	config := mysql.Config{DSN: "cimelli_app:S3c12eT@tcp(127.0.0.1:3306)/cimelli_db"}

	dialect := mysql.New(config)

	Db, err = gorm.Open(dialect, &gorm.Config{Logger: newLogger})

	if err != nil {
		fmt.Println(err.Error())
	}

}
