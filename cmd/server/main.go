package main

import (
	"log"

	"github.com/vominhtrungpro/config"
	charactergenerator "github.com/vominhtrungpro/internal/characters/generator"
	elementgenerator "github.com/vominhtrungpro/internal/elements/generator"
	"github.com/vominhtrungpro/internal/server"
	mysqlserver "github.com/vominhtrungpro/pkg/db/mysql"
	// "gorm.io/driver/mysql"
	// "gorm.io/gen"
	// "gorm.io/gorm"
)

//mysqlserver "github.com/vominhtrungpro/pkg/db/mysql"

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
	s := server.NewServer(
		cfg,
		conn,
	)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
	//generate()
}

// func generate() {
// 	conf := gen.Config{
// 		OutPath: "../../internal/model/dbmodel",
// 		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
// 	}
// 	g := gen.NewGenerator(conf)

// 	gormdb, _ := gorm.Open(mysql.Open("root:tin14091998@tcp(0.0.0.0:3306)/hsr?parseTime=true"))
// 	g.UseDB(gormdb) // reuse your gorm db

// 	// Generate basic type-safe DAO API for struct `model.User` following conventions

// 	g.ApplyBasic(
// 		g.GenerateModel("characters"),
// 	)
// 	g.ApplyBasic(
// 		// Generate structs from all tables of current database
// 		g.GenerateAllTable()...,
// 	)
// 	// Generate the code
// 	g.Execute()
// }
