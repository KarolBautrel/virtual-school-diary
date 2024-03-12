package globalutils

import "os"

func GetDbEnvironmentVariables() (host, user, dbName, password, sslMode, port string) {
	host = os.Getenv("DB_HOST")
	user = os.Getenv("DB_USER")
	dbName = os.Getenv("DB_NAME")
	password = os.Getenv("DB_PASSWORD")
	sslMode = os.Getenv("DB_SSLMODE")
	port = os.Getenv("DB_PORT")
	return
}
