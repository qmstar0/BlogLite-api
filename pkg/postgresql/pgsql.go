package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// GetDB 获取db对象
func GetDB() *gorm.DB {
	return db
}

func InitDB() {
	// dsn := "users:password@tcp(host:port)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	pwd := os.Getenv("DATABASE_PASSWORD_DEV")
	if pwd == "" {
		panic("mysql is not configured: see env:DATABASE_PASSWORD_DEV")
	}
	user := os.Getenv("DATABASE_USER_DEV")
	if user == "" {
		panic("mysql is not configured: see env:DATABASE_USER_DEV")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		"192.168.1.10",
		user,
		pwd,
		"blog_dev",
		5432,
		"disable",
	)
	if err := connectDataBase(dsn); err != nil {
		panic(err)
	}
}

// 连接数据库
func connectDataBase(dsn string) error {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 使用标准输出作为日志输出
			logger.Config{
				SlowThreshold: time.Second, // 慢查询阈值，设为0以便捕获所有查询
				LogLevel:      logger.Info, // 设置日志级别为 Info 或更高级别
				Colorful:      true,        // 在终端中启用彩色输出
			},
		),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return err
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	// 设置连接池大小
	sqlDB.SetMaxOpenConns(100)
	// 设置连接的最大空闲时间
	sqlDB.SetMaxIdleConns(10)
	// 设置连接的最大生存时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 迁移数据库结构
	return nil
}
