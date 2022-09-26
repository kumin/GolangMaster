package provider

import (
	"github.com/google/wire"

	"github.com/kumin/GolangMaster/restful/infras"
	"github.com/kumin/GolangMaster/restful/repos/mysql"
)

var MysqlGraphSet = wire.NewSet(
	infras.InfaGraphSet,
	mysql.NewProductMysqlRepo,
)
