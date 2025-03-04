package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("❌ Gagal memuat file .env")
    } else {
        fmt.Println("✅ .env berhasil dimuat")
    }

    // Ambil konfigurasi dari environment
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Debugging: Cek apakah variabel environment terbaca
    fmt.Println("DB_USER:", dbUser)
    fmt.Println("DB_PASSWORD:", dbPassword)
    fmt.Println("DB_HOST:", dbHost)
    fmt.Println("DB_PORT:", dbPort)
    fmt.Println("DB_NAME:", dbName)

	dsn :=  fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
						dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Println("🟢 Connecting to database with DSN:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
        fmt.Println("❌ Gagal koneksi database:", err)
        return
	}
	// db, err := gorm.Open(mysql.Open("root:@tcp(localhost:5222)/pelatihan_tenis"))
	// if err != nil {
	// 	fmt.Println("Gagal koneksi database")
	// }

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Booking{})
	db.AutoMigrate(&Admin{})

	DB = db
	fmt.Println("✅ Koneksi database berhasil!")
}