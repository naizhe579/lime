package lime

import (
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
	"os"
)

type EnvKey = string

const (
	EnvKeyAddr     EnvKey = "ADDR"
	EnvKeyMysqlDsn EnvKey = "MYSQL_DSN"
)

func InitEnvFile(filePath string) {
	log.Debugf("init env file: %s", filePath)
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
