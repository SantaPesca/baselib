package driver

import (
	"fmt"
	"github.com/Kamva/mgm/v3"
	"github.com/SantaPesca/baselib/pkg/utils"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"time"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", viper.Get("postgres.url"))

	if err != nil {
		utils.MyLog.Fatalf("Cannot connect to Postgres: %v", err)
	}

	db.LogMode(false)

	fmt.Println("Successfully connected to Postgres!")

	return db
}

func ConnectRedisDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.url"),
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
	err := mgm.SetDefaultConfig(nil, "santapesca", options.Client().ApplyURI(viper.GetString("mongo.url")))
	_, client, _, err := mgm.DefaultConfigs()
	err = client.Ping(mgm.Ctx(), nil)
	if err != nil {
		utils.MyLog.Fatalf("Cannot connect to Mongo: %v", err)
	}

	db := client.Database("santapesca")
	indexOpts := options.CreateIndexes().
		SetMaxTime(time.Second * 10)

	// Index to location 2dsphere type
	locationIndexModel := mongo.IndexModel{
		Keys: bsonx.MDoc{"location": bsonx.String("2dsphere")},
	}
	_, err = db.Collection("posts").Indexes().CreateOne(
		mgm.Ctx(),
		locationIndexModel,
		indexOpts,
	)
	if err != nil {
		utils.MyLog.Fatalf("Cannot create 2dsphere index: %v", err)
	}

	// Index to user_id
	userIdIndexModel := mongo.IndexModel{
		Keys: bson.M{"user_id": 1},
	}
	_, err = db.Collection("user_photos").Indexes().CreateOne(
		mgm.Ctx(),
		userIdIndexModel,
		indexOpts,
	)
	if err != nil {
		utils.MyLog.Fatalf("Cannot create user_id index: %v", err)
	}

	fmt.Println("Successfully connected to Mongo!")
}
