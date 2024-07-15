package lime

import (
	"github.com/gofiber/fiber/v3/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB
var onceDB sync.Once

func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	onceDB.Do(func() {
		var err error
		dsn := GetEnvValue(EnvKeyMysqlDsn)
		log.Debugf("connecting to mysql database: %s", dsn)
		db, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			log.Fatal(err)
		}
	})
}
