package app

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/technolingo/bookstore-users-api/utils"
)

var (
	// DB is the database connection
	DB               *gorm.DB
	databaseUsername = os.Getenv("database_username")
	databasePassword = os.Getenv("database_password")
	databaseHostAddr = os.Getenv("database_host_addr")
	databaseHostPort = os.Getenv("database_host_port")
	databaseName     = utils.Getenv("database_name", "bookstore_users")
	databaseSchema   = utils.Getenv("database_schema", "app_data")
	databaseSslmode  = utils.Getenv("database_sslmode", "disabled")
	databaseTimezone = utils.Getenv("database_timezone", "UTC")

	router = gin.Default()
)

// StartApp is the entrypoint of the application
func StartApp() {
	var err error

	// Vanila SQL DB Connection (DB *sql.DB)
	// dbSrcName := fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%s)/%s?charset=utf8",
	// 	databaseUsername,
	// 	databasePassword,
	// 	databaseHostAddr,
	//  databaseHostPort,
	// 	databaseSchema,
	// )
	// DB, err = sql.Open("postgres", dbSrcName)
	// if err != nil {
	// 	panic(err)
	// }
	// if err = DB.Ping(); err != nil {
	// 	panic(err)
	// }

	// DB Connection via Gorm
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		databaseHostAddr,
		databaseUsername,
		databasePassword,
		databaseName,
		databaseHostPort,
		databaseSslmode,
		databaseTimezone,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to the database.")

	mapUrls(router)
	router.Run(":8080")
}
