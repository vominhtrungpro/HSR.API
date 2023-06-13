package main

import (
	"log"

	"github.com/vominhtrungpro/config"
	charactergenerator "github.com/vominhtrungpro/internal/characters/generator"
	elementgenerator "github.com/vominhtrungpro/internal/elements/generator"
	"github.com/vominhtrungpro/internal/server"
	usergenerator "github.com/vominhtrungpro/internal/users/generator"
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
	charactergenerator.InitSnowflakeGenerators()
	elementgenerator.InitSnowflakeGenerators()
	usergenerator.InitSnowflakeGenerators()
	s := server.NewServer(
		cfg,
		conn,
	)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
