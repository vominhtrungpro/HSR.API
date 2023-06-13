package generate

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenerateModel() {
	conf := gen.Config{
		OutPath: "../../internal/model/dbmodel",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	}
	g := gen.NewGenerator(conf)

	gormdb, _ := gorm.Open(mysql.Open("root:tin14091998@tcp(0.0.0.0:3306)/hsr?parseTime=true"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	g.ApplyBasic(
		g.GenerateModel("characters"),
	)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
