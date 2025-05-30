package config

import (
	"context"
	"database/sql"
	"log"
	"os"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var DB *sql.DB
var RDB *redis.Client
var Ctx = context.Background()

func Init() {
	var err error

	DB, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("PostgreSQL Error: ", err)
	}

	RDB = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
}
