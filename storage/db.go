package storage

import (
	"log"

	config "myapp/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB a global for GORM DB object
var DB *gorm.DB

// NewDB to initiate Database Connection
func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

// GetDBInstance return DB object
func GetDBInstance() *gorm.DB {
	return DB
}
func createDBInstanc() *gorm.DB {
	
}
