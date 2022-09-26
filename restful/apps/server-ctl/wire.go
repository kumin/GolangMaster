//go:build wireinject
// +build wireinject

package apps

import (
	"github.com/google/wire"
	"github.com/kumin/GolangMaster/restful/handler"
	"github.com/kumin/GolangMaster/restful/repos/provider"
	"github.com/kumin/GolangMaster/restful/services"
)

var SuperGraphSet = wire.NewSet(
	provider.MysqlGraphSet,
	services.ServiceGraphSet,
	handler.HandlerGraphSet,
	ServerGraphSet,
)

func BuildServer() (*HttpServer, error) {
	wire.Build(
		SuperGraphSet,
	)

	return nil, nil
}
