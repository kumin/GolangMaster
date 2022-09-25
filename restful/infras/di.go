package infras

import "github.com/google/wire"

var InfaGraph = wire.NewSet(
	NewMysqlConnector,
)
