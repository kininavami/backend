package db

import (
	"fmt"
	"github.com/vmware/vending/internal/utils"
	"log"
	"os"
	"gorm.io/gorm/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/driver/postgres"
	"time"
)

// Db gorm db object
var Db *gorm.DB

func InitDB()  {

	var DBReconnectRetry uint64
	var DBReconnectRetryMax uint64 = 5
	var DBReconnectRetryWait uint64 = 10

	db := utils.Getenv("POSTGRES_DB", "postgres")
	du := utils.Getenv("POSTGRES_USER", "postgres")
	dp := utils.Getenv("POSTGRES_PASSWORD", "mysecretpassword")
	dh := utils.Getenv("POSTGRES_HOST", "localhost")


 	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", du, dp, dh, db)
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Info, // Log level

		},
	)

	for {
		log.Printf("DB connection attempt: (%v/%v)", DBReconnectRetry, DBReconnectRetryMax)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{

			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "virtual_vending.", // table name prefix, table for `services` would be `resilience_core.services`
			},
			Logger: dbLogger,
		})

		if err != nil {
			DBReconnectRetry++
			log.Printf("DB Connection Error: %v", err)
			if DBReconnectRetry >= DBReconnectRetryMax {
				log.Fatal(err)
			}
			time.Sleep(time.Duration(DBReconnectRetryWait) * time.Second)
			continue
		}
		log.Printf("Database Connection Succeeded")
		Db = db

		// Set the postgres 'search_path' to 'resilience_core' for all connections
		Db.Exec(`set search_path='virtual_vending'`)
		break
	}
}
