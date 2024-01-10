package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	db *gorm.DB
)

// DBClient 数据库客户端

// GetDB 获取db对象
func GetDB() *gorm.DB {
	return db
}

func init() {

	// dsn := "users:password@tcp(host:port)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	//pwd := os.Getenv("BLOG_MYSQL_PASSWORD")
	//if pwd == "" {
	//	panic("mysql is not configured: see env:BLOG_MYSQL_PASSWORD")
	//}
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
	//	config.Conf.Database.User,
	//	pwd,
	//	config.Conf.Database.Addr,
	//	config.Conf.Database.Name,
	//	config.Conf.Database.Charset,
	//)
	dsn := "host=192.168.1.3 user=qmstar password=@yz021105 dbname=blog_dev port=5432 sslmode=disable"
	if err := connectDataBase(dsn); err != nil {
		panic(err)
	}
}

// 连接数据库
func connectDataBase(dsn string) error {
	var err error
	db_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
	sqlDB, err := db_.DB()
	if err != nil {
		return err
	}
	// 设置连接池大小
	sqlDB.SetMaxOpenConns(100)
	// 设置连接的最大空闲时间
	sqlDB.SetMaxIdleConns(10)
	// 设置连接的最大生存时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	db = db_
	// 迁移数据库结构
	if err = Migrattion(); err != nil {
		return err
	}
	return nil
}
