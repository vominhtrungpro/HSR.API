package generator

import "github.com/vominhtrungpro/pkg/snowflake"

var (
	ProductSNF snowflake.SnowflakeGenerator
)

func InitSnowflakeGenerators() error {
	ProductSNF = snowflake.New()

	return nil
}
