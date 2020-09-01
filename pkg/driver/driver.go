package driver

import (
	"fmt"
	"github.com/Kamva/mgm/v3"
	"github.com/SantaPesca/baselib/pkg/utils"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func ConnectMongoDB() {
	err := mgm.SetDefaultConfig(nil, "santapesca", options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	_, client, _, err := mgm.DefaultConfigs()
	err = client.Ping(mgm.Ctx(), nil)
	if err != nil {
		utils.MyLog.Fatalf("Cannot connect to Mongo: %v", err)
	}
	fmt.Println("Successfully connected to Mongo!")
}
