package mysql

import "testing"

func TestOpenMysql(t *testing.T) {
	OpenMysql(&Config{
		Host:     "118.24.39.34:3306",
		User:     "user",
		Password: "bBFt4aEzcWkKrhHD",
		Database: "user",
	})
}
