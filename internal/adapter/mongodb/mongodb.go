package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const (
	database = "blog_dev"
)

var (
	db *mongo.Database
)

func GetDB() *mongo.Database {
	if db == nil {
		panic("database(mongodb) not init")
	}
	return db
}

func Init() (closeFn func(context.Context) error) {

	host := os.Getenv("DATABASE_HOST_DEV")
	if host == "" {
		panic("database is not configured: see env:DATABASE_HOST_DEV")
	}
	pwd := os.Getenv("DATABASE_PASSWORD_DEV")
	if pwd == "" {
		panic("database is not configured: see env:DATABASE_PASSWORD_DEV")
	}
	user := os.Getenv("DATABASE_USER_DEV")
	if user == "" {
		panic("database is not configured: see env:DATABASE_USER_DEV")
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:2717/?authSource=user&connect=direct", user, pwd, host))

	ctx := context.TODO()
	// 连接到 MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	// 检查连接是否成功
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	db = client.Database(database)

	// 关闭连接的方法
	return client.Disconnect
}
