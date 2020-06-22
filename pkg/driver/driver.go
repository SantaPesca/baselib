package driver

import (
	"fmt"
	"github.com/SantaPesca/baselib/pkg/utils"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"os"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		utils.MyLog.Fatalf("Cannot connect to Postgres: %v", err)
	}

	db.LogMode(false)

	fmt.Println("Successfully connected to Postgres!")

	return db
}

func ConnectRedisDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Ping().Err()
	if err != nil {
		utils.MyLog.Fatalf("Cannot connect to Redis: %v", err)
	}

	fmt.Println("Successfully connected to Redis!")

	return rdb
}
