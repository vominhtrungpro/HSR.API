package main

import (
	"log"

	"github.com/vominhtrungpro/config"
	charactergenerator "github.com/vominhtrungpro/internal/characters/generator"
	elementgenerator "github.com/vominhtrungpro/internal/elements/generator"
	"github.com/vominhtrungpro/internal/server"
	usergenerator "github.com/vominhtrungpro/internal/users/generator"
	redisclient "github.com/vominhtrungpro/pkg/cache/redis"
	mysqlserver "github.com/vominhtrungpro/pkg/db/mysql"
)

func main() {
	log.Println("Starting api server")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	// Connect to database
	conn, err := mysqlserver.New(&cfg.MySQL)
	if err != nil {
		log.Fatal(err)
	}
	rdb, err := redisclient.NewClient2(&cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}
	charactergenerator.InitSnowflakeGenerators()
	elementgenerator.InitSnowflakeGenerators()
	usergenerator.InitSnowflakeGenerators()
	s := server.NewServer(
		cfg,
		conn,
		server.Redis(rdb),
	)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}

}

// func ExampleClient() {
// 	ctx := context.Background()
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})

// 	err := rdb.Set(ctx, "key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := rdb.Get(ctx, "mykey").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := rdb.Get(ctx, "key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}

// 	val3, err := rdb.Sort(ctx, "list", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val3)
// }
