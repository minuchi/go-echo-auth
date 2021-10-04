package database

import (
	"fmt"
	"github.com/minuchi/go-echo-auth/lib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

func getDatabaseConfigString(databaseConfig lib.DatabaseConfig) string {
	databaseConfigMap := map[string]string{
		"host":     databaseConfig.Host,
		"user":     databaseConfig.User,
		"password": databaseConfig.Password,
		"dbname":   databaseConfig.DbName,
		"port":     databaseConfig.Port,
		"sslmode":  databaseConfig.Sslmode,
		"TimeZone": databaseConfig.Timezone,
	}

	var dns string
	for key, value := range databaseConfigMap {
		dns += fmt.Sprintf("%s=%s ", key, value)
	}
	return strings.TrimSpace(dns)
}

func Connect(databaseConfig lib.DatabaseConfig) *gorm.DB {
	dsn := getDatabaseConfigString(databaseConfig)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
