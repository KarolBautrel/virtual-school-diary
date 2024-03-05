package envUtils

import "os"

func GetDbEnvironmentVariables() (host, user, dbName, password, sslMode string) {
	host = os.Getenv("DB_HOST")
	user = os.Getenv("DB_USER")
	dbName = os.Getenv("DB_NAME")
	password = os.Getenv("DB_PASSWORD")
	sslMode = os.Getenv("DB_SSLMODE")
	return
}
