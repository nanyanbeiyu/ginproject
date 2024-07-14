/*
@author: NanYan
*/
package snowflakeID

import (
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

func SnowflakeID() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		zap.L().Error("snowflake.NewNode(1) error", zap.Error(err))
		return 0
	}
	id := node.Generate()
	return id.Int64()
}
