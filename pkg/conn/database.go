package conn

import (
	"initial_project_go/internal/user"
	"initial_project_go/pkg/config"
	"log"
	"time"

	// "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const TIMECONN = 10

var DATABASE *gorm.DB

// ------------------> jika menggunakan mysql <------------------
// func DatabaseConn() {
// 	var err error
// 	dsn := config.GetConfig("database.user") + ":" + config.GetConfig("database.pass") + "@tcp(" + config.GetConfig("database.host") + ":" + config.GetConfig("database.port") + ")/" + config.GetConfig("database.name") + "?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	sqlDb, err := db.DB()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	sqlDb.SetMaxIdleConns(5)
// 	sqlDb.SetMaxOpenConns(100)
// 	sqlDb.SetConnMaxLifetime(time.Minute * TIMECONN)
// 	DATABASE = db
// }

// ------------------> jika menggunakan postgre <------------------
func DatabaseConn() {
	var err error
	// PostgreSQL DSN format
	dsn := "host=" + config.GetConfig("database.host") +
		" user=" + config.GetConfig("database.user") +
		" password=" + config.GetConfig("database.pass") +
		" dbname=" + config.GetConfig("database.name") +
		" port=" + config.GetConfig("database.port") +
		" sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err)
		return
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Println(err)
		return
	}
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Minute * TIMECONN)
	DATABASE = db

	// Panggil fungsi migrasi setelah koneksi berhasil
	DatabaseMigrate()
}

func DatabaseMigrate() {
	// Pastikan koneksi ke database telah dibuat
	if DATABASE == nil {
		log.Println("Database connection is not initialized")
		return
	}

	// Migrasi model yang diinginkan
	err := DATABASE.AutoMigrate(&user.Users{})
	if err != nil {
		log.Println("Error migrating database:", err)
		return
	}

	log.Println("Database migrated successfully")
}
