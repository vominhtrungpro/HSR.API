package snowflake

import (
	pkgerrors "github.com/pkg/errors"
)

// Generate generates Sonyflake ID
func (s SnowflakeGenerator) Generate() (int64, error) {
	n, err := s.snowflake.NextID()
	if err != nil {
		return 0, pkgerrors.WithStack(err)
	}

	return int64(n), nil
}
