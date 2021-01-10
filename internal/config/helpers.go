package config

import (
	"fmt"
	"github.com/authenter/nyx/pkg/db"
)

// GetAddr returns a string with the combination from server host and port
func (sc *serverConfig) GetAddr() string {
	return fmt.Sprintf("%s:%d", sc.Host, sc.Port)
}

// GetDSN return the data source name string for the matching database driver
func (dc *databaseConfig) GetDSN() string {
	switch dc.Driver {
	case db.DriverMySQL:
		return fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?multiStatements=true&parseTime=true&loc=UTC&collation=utf8mb4_general_ci",
			dc.Username,
			dc.Password,
			dc.Host,
			dc.Port,
			dc.Name,
		)
	}

	return ""
}
