package config

import "fmt"

const (
	DBHost = "serverdb"
	DBPort = "3306"
	DBUser = "mysql"
	DBPass = "mysql"
	DBName = "mydatbase"
)

func GetDBConnection() string {

	connection := fmt.
		Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			DBUser,
			DBPass,
			DBHost,
			DBPort,
			DBName)
	return connection
}
